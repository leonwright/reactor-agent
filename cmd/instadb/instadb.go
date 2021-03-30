package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leonwright/devhelper/pkg/docker"
	"github.com/leonwright/devhelper/pkg/generator"
	"github.com/urfave/cli/v2"
)

func main() {
	var databaseName string

	app := &cli.App{
		Name:  "instadb",
		Usage: "instantly create development databases!",
	}

	nameFlag := &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "the name for the database",

		Destination: &databaseName,
	}

	app.Commands = []*cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a new database for development",
			Subcommands: []*cli.Command{
				{
					Name:  "mongo",
					Usage: "create a mongo development database",
					Action: func(c *cli.Context) error {
						if databaseName == "" {
							databaseName = generator.GenerateCodeName()
						}
						err := docker.CreateDB(databaseName, "mongo")
						if err != nil {
							fmt.Errorf("There was an error creating the database: %v", err)
						}
						return nil
					},
					Flags: []cli.Flag{
						nameFlag,
					},
				},
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete a database",
			Action: func(c *cli.Context) error {
				if databaseName == "" {
					databaseName = generator.GenerateCodeName()
				}
				fmt.Println("deleting:", databaseName)
				return nil
			},
			Flags: []cli.Flag{
				nameFlag,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
