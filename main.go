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

		if c.NArg() > 0 {
			fmt.Print("This app does not accept arguments!\n")

			return nil
		}

		if delimiter == "." || delimiter == "-" || delimiter == "_" || delimiter == " " {
			delim = delimiter
		} else if len(delimiter) > 0 {
			fmt.Print("That delimiter is not allowed! Using the default.\n")
		}

		name := Generate(delim)

		fmt.Printf("%s\n", name)

		return nil
	}

	app.Run(os.Args)
}

// Generate produces a name using a random adjective and noun.
// delim is added as a delimiter between words.
func Generate(delim string) string {
	adj := Adjectives[rand.Intn(len(Adjectives))]
	noun := Colours[rand.Intn(len(Colours))]

	return fmt.Sprintf("%s%s%s", adj, delim, noun)

}
