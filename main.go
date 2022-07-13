package main

import (
	"fmt"
	"log"

	"github.com/s-vvardenfell/goropes/rope"
)

const ( //not work with non-ascii yet
	// stropeRu = "Строп (канат, корд) — структура данных, которая позволяет эффективно хранить и обрабатывать длинные строки, например текст."
	stropeEn = "A sling (rope, cord) is a data structure that allows you to efficiently store and process long strings, such as text."
)

func main() {
	r, err := rope.NewRopeFromString(stropeEn, 8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
