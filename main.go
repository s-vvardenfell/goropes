package main

import (
	"fmt"
	"goropes/goropes"
)

func main() {
	s := "some vague string"
	r := goropes.NewRopeFromString(s)
	fmt.Println(r)
}
