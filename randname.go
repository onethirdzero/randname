package randname

import (
	"fmt"
	"math/rand"
	"time"
)

// This is in a separate package from main.go because this is meant to be used
// as a library, while main.go is meant to built as a binary.
// https://stackoverflow.com/questions/14284375/can-i-have-a-library-and-binary-with-the-same-name

// GenerateLowerCase produces a lower case name using a random adjective and noun.
// delim is added as a delimiter between words.
func GenerateLowerCase(delim string) string {
	rand.Seed(time.Now().UTC().UnixNano())

	adj := Adjectives[rand.Intn(len(Adjectives))]
	noun := Colours[rand.Intn(len(Colours))]

	return fmt.Sprintf("%s%s%s", adj, delim, noun)

}

// GenerateCamelCase produces a camel case name using a random adjective and noun.
// delim is added as a delimiter between words.
func GenerateCamelCase(delim string) string {
	rand.Seed(time.Now().UTC().UnixNano())

	adj := Adjectives[rand.Intn(len(Adjectives))]
	noun := Colours[rand.Intn(len(Colours))]

	titledAdj := strings.Title(string(adj))
	titledNoun := strings.Title(string(noun))

	return fmt.Sprintf("%s%s%s", titledAdj, titledNoun, noun)

}
