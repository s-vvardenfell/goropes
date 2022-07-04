package main

import (
	"fmt"

	"github.com/s-vvardenfell/goropes/rope"
)

const (
	strope = "Строп (канат, корд) — структура данных, которая позволяет эффективно хранить и обрабатывать длинные строки, например текст."
)

func main() {
	// f, err := os.Open("temp/test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// r := rope.NewRopeFromReader(f)
	r := rope.NewRopeFromString(strope)

	fmt.Println(r.Root, r.Root.Left, r.Root.Right)
}
