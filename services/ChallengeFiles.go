package services

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ruiblaese/data-integration-challenge/handlers"

	"github.com/coscms/xorm"
)

//ProcessesChallengeFiles processa arquivos do desafio
func ProcessesChallengeFiles() (string, error) {

	filePackage := "./assets/dataIntegrationChallenge.tgz"
	if _, err := os.Stat(filePackage); !os.IsNotExist(err) {

		folderTempExtract := "./temp/filesChallenge/"

		if _, err := os.Stat(folderTempExtract); !os.IsNotExist(err) {
			os.RemoveAll(folderTempExtract)
		}

		if err := ExtractFiles(filePackage, folderTempExtract); err == nil {

			return folderTempExtract, nil

		} else {

			return "", errors.New("ProcessesChallengeFiles: extractFile: " + filePackage + " error:" + fmt.Sprint(err))
		}

	} else {
		return "", errors.New("ProcessesChallengeFiles: file not found: " + filePackage)
	}
}

//ProcessesFirstData processa arquivos do desafio
func ProcessesFirstData(xormEngine *xorm.Engine) error {

	if folder, err := ProcessesChallengeFiles(); err == nil {

		fileFirstData := folder + "q1_catalog.csv"
		handlers.RegisterCompanyFromCSV(xormEngine, fileFirstData)

		//handlers.RegisterCompanyFromCSV(xormEngine, folder+"q2_clientData.csv")

		if _, err := os.Stat(folder); !os.IsNotExist(err) {
			os.RemoveAll(folder)
		}

	} else {
		log.Fatalln(err)
	}
	return nil
}
