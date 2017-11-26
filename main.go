package main

import (
	"fmt"
	"math/rand"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Name = "randnamegen"
	app.Usage = "Generate random names"

	// The CLI app's main function.
	app.Action = func(c *cli.Context) error {
		adj := Adjectives[rand.Intn(len(Adjectives))]
		noun := Colours[rand.Intn(len(Colours))]

		name := fmt.Sprintf("%s-%s", adj, noun)

		fmt.Printf("%s\n", name)

		return nil
	}

	app.Run(os.Args)
}
