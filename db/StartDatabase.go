package db

import (
	"github.com/coscms/xorm"
	//sqlite3
	_ "github.com/mattn/go-sqlite3"
	"github.com/ruiblaese/data-integration-challenge/models"
)

//StartDatabase inicia banco de dados -> verifica se existe, se nao existe cria
func StartDatabase() (*xorm.Engine, error) {

	engine, err := xorm.NewEngine("sqlite3", "./challenge.db")

	engine.ShowSQL(false)
	engine.ShowExecTime(false)

	if err != nil {
		return nil, err
	}
	err = engine.Sync(new(models.Company))

	if err != nil {
		return nil, err
	}

	return engine, nil
}
