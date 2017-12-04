package main

import (
	"fmt"
	"os"

	"github.com/onethirdzero/randname"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Name = "randname"
	app.Usage = "Generate random names"

	// Flags.
	var delimiter string
	var camelCase bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "delimiter, d",
			Value:       "",
			Usage:       "Delimiter used to separate the name words",
			Destination: &delimiter,
		},
		cli.BoolFlag{
			Name:        "camelcase, c",
			Usage:       "Return a camel case name instead of a lower case name",
			Destination: &camelCase,
		},
	}

	// The CLI app's main function.
	app.Action = func(c *cli.Context) error {
		var delim string
		var name string

		if c.NArg() > 0 {
			fmt.Print("This app does not accept arguments!\n")

			return nil
		}

		if delimiter == "." || delimiter == "-" || delimiter == "_" || delimiter == " " {
			delim = delimiter
		} else if len(delimiter) > 0 {
			fmt.Print("That delimiter is not allowed! Using the default.\n")
		}

		if camelCase {
			name = randname.GenerateCamelCase(delim)
		} else {
			name = randname.GenerateLowerCase(delim)
		}

		fmt.Printf("%s\n", name)

		return nil
	}

	app.Run(os.Args)
}
