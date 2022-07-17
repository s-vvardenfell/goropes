package main

import (
	"log"

	rope "github.com/s-vvardenfell/goropes/rope"
)

const (
	stropeRu = "Строп (канат, корд) — структура данных, которая позволяет эффективно хранить и обрабатывать длинные строки, например текст."
	stropeEn = "A sling (rope, cord) is a data structure that allows you to efficiently store and process long strings, such as text."
)

func main() {
	r1, err := rope.NewGoRopeFromString(stropeRu, 8)
	if err != nil {
		log.Fatal(err)
	}
	r1.Display()

	r2, err := rope.NewGoRopeFromString(stropeEn, 8)
	if err != nil {
		log.Fatal(err)
	}
	r2.Display()

	r := r1.Concat(r2)

	r.Display()
}
