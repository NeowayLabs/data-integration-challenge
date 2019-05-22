package actions

import (
	"log"

	"github.com/jean-lopes/data-integration-challenge/pkg/httphelpers"
	"github.com/jean-lopes/data-integration-challenge/pkg/models"
	"github.com/jean-lopes/data-integration-challenge/pkg/util"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli"
)

// Create a company
func Create(c *cli.Context) {
	id := readID(c)
	ws := c.String("website")

	company := models.Company{
		ID:      id,
		Name:    c.String("name"),
		Zip:     c.String("zip"),
		Website: &ws,
	}

	url := c.GlobalString("base-url") + "/companies"
	response, err := httphelpers.Post(url, company)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(string(response))
	}
}

func readID(c *cli.Context) *uuid.UUID {
	idStr := c.String("id")

	if util.IsBlank(idStr) {
		return nil
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &id
}
