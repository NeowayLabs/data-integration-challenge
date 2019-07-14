package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ruiblaese/data-integration-challenge/services"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ruiblaese/data-integration-challenge/db"
	"github.com/ruiblaese/data-integration-challenge/routes"
)

func main() {

	processFirstData := false
	if _, err := os.Stat("./challenge.db"); os.IsNotExist(err) {
		processFirstData = true
	}

	xormEngine, err := db.StartDatabase()
	if err != nil {
		log.Fatalln("Error in StartDatabase->", err)
	}
	defer xormEngine.Close()

	if processFirstData {
		fmt.Println("Processes first data!")
		services.ProcessesFirstData(xormEngine)
	}

	ginRouter := gin.Default()
	ginRouter = routes.StartRouter(ginRouter, xormEngine)

	ginRouter.Run(":4000")

}
