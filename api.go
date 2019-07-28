package main

import(
    "github.com/gorilla/mux"
    "database/sql"
    "encoding/csv"
    "encoding/json"
    "io"
    "strconv"
    "net/http"
    "io/ioutil"
    "github.com/joho/godotenv"
    "log"
    "strings"
    "os"
    "fmt"
    _ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() {
        
    e := godotenv.Load()
    if e != nil {
        fmt.Print(e)
    }

    username := os.Getenv("db_user")
    password := os.Getenv("db_pass")
    dbName := os.Getenv("db_name")
    dbHost := os.Getenv("db_host")
    dbPort := os.Getenv("db_port")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s "+
        "sslmode=disable", dbHost, dbPort, username, password, dbName)

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
}

type Company struct {
    Id    string `json:"Id,omitempty"`
    Name  string `json:"Name,omitempty"`
    Zip  string `json:"Zip,omitempty"`
    Website  string `json:"Website,omitempty"`
}

var companies []Company

func main() {

    newRouter := mux.NewRouter().StrictSlash(true)
    newRouter.HandleFunc("/", index)
    newRouter.HandleFunc("/company", GetCompanies).Methods("GET")
    newRouter.HandleFunc("/company/{id}", GetCompany).Methods("GET")
    newRouter.HandleFunc("/company/match", MatchZipAndName).Methods("POST")
    newRouter.HandleFunc("/populate", PopulateDB).Methods("POST")
    newRouter.HandleFunc("/integrate_website", IntegrateWebsite).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", newRouter))

}

func index(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/company", http.StatusSeeOther)
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(companies)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    for _, item := range companies {
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Company{})
}

func MatchZipAndName(w http.ResponseWriter, r *http.Request) {

     var match_company Company

     err := json.NewDecoder(r.Body).Decode(&match_company)
     if err != nil {
        panic(err)
     }

     for _, item := range companies {
         if (strings.Contains(item.Name, match_company.Name) && (item.Zip == match_company.Zip)) {
            json.NewEncoder(w).Encode(item)
            return
         }
    }
    json.NewEncoder(w).Encode(&Company{})
}

func PopulateDB(w http.ResponseWriter, r *http.Request) {

    reqBody, _ := ioutil.ReadAll(r.Body)    

    delete_query := `DROP TABLE IF EXISTS companies;`

    _, err = db.Exec(delete_query)
    if err != nil {
        panic(err)
    }

    create_query := `
    CREATE TABLE IF NOT EXISTS companies (
            Id serial primary key,
            name varchar,
            zip varchar
    );`

    _, err = db.Exec(create_query)
    if err != nil {
        panic(err)
    }

    insert_query := `
    INSERT INTO companies (Id, name, zip)
    VALUES ($1, $2, $3);`

    csvFile, _ := os.Open(string(reqBody))
    reader := csv.NewReader(csvFile)
    reader.Comma = ';'
    id := -1 
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        } else if id == -1 {
            id += 1
            continue
        }

        id += 1

        name := strings.ToUpper(line[0])
        zip := fmt.Sprintf("%05s", line[1])

        companies = append(companies, Company{Id: strconv.Itoa(id), Name: name, Zip: zip})

        _, err = db.Exec(insert_query, id, name, zip)
        if err != nil {
            panic(err)
        }   
    }
}


func IntegrateWebsite(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)    

    delete_query := `DROP TABLE IF EXISTS aux_website;`

    _, err = db.Exec(delete_query)
    if err != nil {
        panic(err)
    }

    create_query := `
    CREATE TABLE IF NOT EXISTS aux_website (
            aux_name varchar primary key,
            aux_zip varchar,
            website varchar
    );`

    _, err = db.Exec(create_query)
    if err != nil {
        panic(err)
    }

    insert_query := `
    INSERT INTO aux_website (aux_name, aux_zip, website)
    VALUES ($1, $2, $3);`

    csvFile, _ := os.Open(string(reqBody))
    reader := csv.NewReader(csvFile)
    reader.Comma = ';'
    id := -1 
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        } else if id == -1 {
            id += 1
            continue
        }

        id += 1

        name := strings.ToUpper(line[0])
        zip := fmt.Sprintf("%05s", line[1])
        website := strings.ToLower(line[2])

        slice_id := 0
        for _, item := range companies {
            slice_id += 1
            if item.Name == name {
                companies = append(companies[:slice_id-1], companies[slice_id:]...)
                companies = append(companies, Company{Id: strconv.Itoa(slice_id), Name: name, Zip: zip, Website: website})
                break
            }
        }

        _, err = db.Exec(insert_query, name, zip, website)
        if err != nil {
            panic(err)
        }
    }

    delete_temp_query := `DROP TABLE IF EXISTS temp_table;`

    _, err = db.Exec(delete_temp_query)
    if err != nil {
        panic(err)
    }

    join_query := `
    CREATE TABLE IF NOT EXISTS temp_table AS
        SELECT companies.*, aux_website.website
        FROM companies
        LEFT JOIN aux_website ON companies.name = aux_website.aux_name;`

    _, err = db.Exec(join_query)
    if err != nil {
        panic(err)
    }

    delete_main_query := `
    DROP TABLE IF EXISTS companies;`

    _, err = db.Exec(delete_main_query)
    if err != nil {
        panic(err)
    }

    set_temp_to_main_query := `
    CREATE TABLE IF NOT EXISTS companies AS
        SELECT * FROM temp_table;`

    _, err = db.Exec(set_temp_to_main_query)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec(delete_temp_query)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec(delete_query)
    if err != nil {
        panic(err)
    }

}
