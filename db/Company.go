package db

import (
	"errors"
	"strconv"
	"strings"

	"github.com/coscms/xorm"
	"github.com/ruiblaese/data-integration-challenge/models"
	//"github.com/ruiblaese/data-integration-challenge/services"
)

//FindCompany retorna todas empresas
func FindCompany(xormEngine *xorm.Engine, id int, name, zipcode string, partialName bool, limit int) ([]models.Company, error) {

	where := "(compan_id is not null)"

	if id != 0 {
		where = where + " and (compan_id = " + strconv.Itoa(id) + ")"
	}
	if name != "" {
		if partialName {
			where = where + " and (compan_name like '" + strings.ToUpper(name) + "%')"
		} else {
			where = where + " and (compan_name = '" + strings.ToUpper(name) + "')"
		}
	}
	if zipcode != "" {
		where = where + " and (compan_zipcode = '" + zipcode + "')"
	}

	var companys []models.Company
	err := xormEngine.Where(where).Limit(limit, 0).Find(&companys)
	if err != nil {
		return []models.Company{}, err
	}
	return companys, nil

}

//InsertCompany cadastra empresa e retorna com id
func InsertCompany(xormEngine *xorm.Engine, company models.Company) (models.Company, error) {

	if company, err := normalizeCompany(company); err == nil {

		id, err := xormEngine.Insert(company)
		if err != nil {
			return models.Company{}, err
		}
		company.ID = int(id)
		return company, nil

	} else {
		return models.Company{}, err
	}

}

//UpdateCompanyByID atualiza empresa
func UpdateCompanyByID(xormEngine *xorm.Engine, id int, company models.Company) (models.Company, error) {

	if company, err := normalizeCompany(company); err == nil {

		_, err := xormEngine.Id(id).Update(&company)
		if err != nil {
			return models.Company{}, err
		}
		return company, nil

	} else {
		return models.Company{}, err
	}
}

//DeleteCompanyByID atualiza empresa
func DeleteCompanyByID(xormEngine *xorm.Engine, id int) (bool, error) {

	_, err := xormEngine.Id(id).Delete(&models.Company{})

	return err == nil, err

}

func validCompany(company models.Company) error {

	var listErrors []string

	if company.Name == "" {
		listErrors = append(listErrors, "Name empty")
	}

	if company.Zipcode == "" {
		listErrors = append(listErrors, "Zipcode empty")
	}

	if len(listErrors) > 0 {
		return errors.New(strings.Join(listErrors, " "))
	}
	return nil
}

func normalizeCompany(company models.Company) (models.Company, error) {

	if err := validCompany(company); err == nil {

		company.Name = strings.ToUpper(company.Name)
		company.Zipcode = Lpad(company.Zipcode, 5, "0")
		if company.Website != "" {
			company.Website = strings.ToLower(company.Website)
		}
		return company, nil

	} else {
		return company, err
	}
}

// Lpad funcao para adicionar caracteres a uma string
// Nao sei por que nao consegui colocar em services, referencia cincurlar :-/
func Lpad(s string, plength int, pad string) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}
