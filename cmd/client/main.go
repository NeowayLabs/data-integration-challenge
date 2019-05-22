package main

import (
	"log"
	"os"

	"github.com/jean-lopes/data-integration-challenge/cmd/client/actions"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Data integration challenge client"
	app.Usage = "Yes" // kkkkk
	app.Author = "Jean Lopes"
	app.Version = "pre-alpha"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "base-url"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Create a company",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "id"},
				cli.StringFlag{Name: "name"},
				cli.StringFlag{Name: "zip"},
				cli.StringFlag{Name: "website"},
			},
			Action: actions.Create,
		},
		{
			Name:  "merge-website",
			Usage: "Merge website for a company",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "name"},
				cli.StringFlag{Name: "zip"},
				cli.StringFlag{Name: "website"},
			},
			Action: actions.MergeWebsite,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
