package rope2

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type stringTuple struct {
	left  string
	right string
}

// Returns visual of Rope tree
func (r *GoRope) Display() {
	lines, _, _, _ := r.displayAux(r.root)
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Returns list of strings, width, height and horizontal coordinate of the root
func (r *GoRope) displayAux(root *Node) ([]string, int, int, int) {
	// no child case
	if root != nil && root.Data != nil {
		line := visualizer(root)
		width := utf8.RuneCountInString(line)
		height := 1
		middle := width / 2
		return []string{line}, width, height, middle
	}

	left, n, p, x := r.displayAux(root.Left)
	rigth, m, q, y := r.displayAux(root.Right)
	s := visualizer(root)
	u := len(s)

	firstLine := strings.Repeat(" ", x+1) + strings.Repeat("_", n-x-1) + s + strings.Repeat("_", y) + strings.Repeat(" ", m-y)
	secondLine := strings.Repeat(" ", x) + "/" + strings.Repeat(" ", n-x-1+u+y) + "\\" + strings.Repeat(" ", m-y-1)

	if p < q {
		for i := 0; i < q-p; i++ {
			left = append(left, strings.Repeat(" ", n))
		}
	} else if q < p {
		for i := 0; i < p-q; i++ {
			rigth = append(rigth, strings.Repeat(" ", m))
		}
	}

	zippedLines := zip(left, rigth)
	lines := []string{firstLine, secondLine}

	for i := range zippedLines {
		lines = append(lines,
			zippedLines[i].left+strings.Repeat(" ", u)+zippedLines[i].right)
	}

	return lines, n + m + u, max(p, q) + 2, n + u/2
}

func visualizer(n *Node) string {
	if n.Data != nil {
		return string(n.Data)
	}
	return strconv.Itoa(n.Weight)
}

func zip(a []string, b []string) []stringTuple /*, error)*/ {
	// if len(a) != len(b) {
	// 	return nil, fmt.Errorf("zip: arguments length must be same ")
	// }

	var t []stringTuple
	for index, value := range a {
		t = append(t, stringTuple{value, b[index]})
	}
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
