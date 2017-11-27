package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	app := cli.NewApp()

	app.Name = "randname"
	app.Usage = "Generate random names"

	// Flags.
	var delimiter string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "delimiter, d",
			Value:       "",
			Usage:       "Delimiter used to separate the name words ",
			Destination: &delimiter,
		},
	}

	// The CLI app's main function.
	app.Action = func(c *cli.Context) error {
		var delim string
		var adj string
		var noun string

		if c.NArg() > 0 {
			fmt.Print("This app does not accept arguments!\n")

			return nil
		}

		if delimiter == "." || delimiter == "-" || delimiter == "_" || delimiter == " " {
			delim = delimiter
		} else if len(delimiter) > 0 {
			fmt.Print("That delimiter is not allowed! Using the default.\n")
		}

		adj = Adjectives[rand.Intn(len(Adjectives))]
		noun = Colours[rand.Intn(len(Colours))]

		name := fmt.Sprintf("%s%s%s", adj, delim, noun)

		fmt.Printf("%s\n", name)

		return nil
	}

	app.Run(os.Args)
}
