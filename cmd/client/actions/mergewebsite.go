package actions

import (
	"log"

	"github.com/jean-lopes/data-integration-challenge/pkg/handlers"
	"github.com/jean-lopes/data-integration-challenge/pkg/httphelpers"
	"github.com/urfave/cli"
)

// MergeWebsite using the HTTP API
func MergeWebsite(c *cli.Context) {
	data := handlers.MergeCompanyWebsiteHandlerBody{
		Name:    c.String("name"),
		Zip:     c.String("zip"),
		Website: c.String("website"),
	}

	url := c.GlobalString("base-url") + "/merge-company-website"
	response, err := httphelpers.Post(url, data)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(string(response))
	}
}
