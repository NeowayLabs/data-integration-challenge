package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/coscms/xorm"
	"github.com/gin-gonic/gin"
	"github.com/ruiblaese/data-integration-challenge/db"
	"github.com/ruiblaese/data-integration-challenge/models"
)

// GetCompany retorna todos das emnpresas
func GetCompany(c *gin.Context) {
	ixorm, _ := c.Get("xorm")
	xormEngine, _ := ixorm.(*xorm.Engine)

	id, _ := strconv.Atoi(c.Query("companyId"))
	name := c.Query("name")
	zipcode := c.Query("zipcode")

	listCompany, _ := db.FindCompany(xormEngine, id, name, zipcode, true, 0)

	c.Header("Content-Type", "application/json")
	if len(listCompany) > 0 {
		c.JSON(http.StatusOK, listCompany)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}

// GetCompanyByID retornoa empresa por id
func GetCompanyByID(c *gin.Context) {
	ixorm, _ := c.Get("xorm")
	xormEngine, _ := ixorm.(*xorm.Engine)
	c.Header("Content-Type", "application/json")

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		if companies, err := db.FindCompany(xormEngine, id, "", "", true, 1); err == nil && len(companies) > 0 {
			c.JSON(http.StatusOK, companies[0])
		} else {
			if err == nil {
				c.AbortWithStatus(http.StatusNotFound)
			} else {
				c.JSON(http.StatusInternalServerError,
					gin.H{"code": "ERROR", "message": "Internal Server Error "})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": "ERROR", "message": "Invalid param"})
	}

}

// PutCompany atualiza informacoes do usuario
func PutCompany(c *gin.Context) {
	ixorm, _ := c.Get("xorm")
	xormEngine, _ := ixorm.(*xorm.Engine)

	c.Header("Content-Type", "application/json")
	var company models.Company
	c.Bind(&company)

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		if companyUpdated, err := db.UpdateCompanyByID(xormEngine, id, company); err == nil && companyUpdated.ID > 0 {
			c.JSON(http.StatusOK, companyUpdated)
		} else {
			if err == nil {
				c.AbortWithStatus(http.StatusNotModified)
			} else {
				c.JSON(http.StatusInternalServerError,
					gin.H{"code": "ERROR", "message": "Internal Server Error "})
			}
		}

	} else {
		c.JSON(http.StatusNoContent, gin.H{"code": "ERROR", "message": "Invalid param"})
	}
}

// NewCompany cria novo usuario
func NewCompany(c *gin.Context) {
	ixorm, _ := c.Get("xorm")
	xormEngine, _ := ixorm.(*xorm.Engine)

	c.Header("Content-Type", "application/json")

	var company models.Company
	c.Bind(&company)

	if companyInserted, err := db.InsertCompany(xormEngine, company); err == nil && companyInserted.ID > 0 {
		c.JSON(http.StatusCreated, companyInserted)

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
	}
}

// DeleteCompany deleta usuario
func DeleteCompany(c *gin.Context) {
	ixorm, _ := c.Get("xorm")
	xormEngine, _ := ixorm.(*xorm.Engine)

	c.Header("Content-Type", "application/json")

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		if apagou, err := db.DeleteCompanyByID(xormEngine, id); apagou && err == nil {
			c.JSON(http.StatusOK, gin.H{"code": "OK", "message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
		}

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR", "message": "Internal Server Error"})
	}
}

// UploadCompanysWithCSV atualiza informacoes do usuario
func UploadCompanysWithCSV(c *gin.Context) {
	ixorm, _ := c.Get("xorm")
	xormEngine, _ := ixorm.(*xorm.Engine)

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	fileTemp := "./temp/" + file.Filename
	c.SaveUploadedFile(file, fileTemp)

	RegisterCompanyFromCSV(xormEngine, fileTemp)

	os.Remove(fileTemp)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

//RegisterCompanyFromCSV cadastra empresas por arquivo CSV
func RegisterCompanyFromCSV(xormEngine *xorm.Engine, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	lines, err := reader.ReadAll()
	if err != nil {
		log.Panicln(err)
	}

	for i, line := range lines {

		if i != 0 {

			name := line[0]
			zipcode := line[1]
			website := ""

			if len(line) > 2 {
				website = line[2]
			}

			company := models.Company{
				Name:    name,
				Zipcode: zipcode,
				Website: website,
			}

			if companiesFinded, err := db.FindCompany(xormEngine, 0, name, "", false, 1); len(companiesFinded) > 0 && err == nil {
				company.ID = companiesFinded[0].ID
				company.Version = companiesFinded[0].Version
				db.UpdateCompanyByID(xormEngine, company.ID, company)
			} else {
				db.InsertCompany(xormEngine, company)
			}
		}

	}
}
