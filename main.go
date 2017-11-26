package main

import (
	"fmt"
	"math/rand"
)

func main() {
	adj := Adjectives[rand.Intn(len(Adjectives))]
	noun := Colours[rand.Intn(len(Colours))]

	name := fmt.Sprintf("%s-%s", adj, noun)

	fmt.Printf("%s\n", name)
}
