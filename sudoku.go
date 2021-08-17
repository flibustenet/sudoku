package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

//go:embed empty.json
var empty_json []byte

type Board [][]int

func (b Board) Set(x, y, v int) {
	b[y][x] = v
}
func (b Board) Get(x, y int) int {
	return b[y][x]
}
func (b Board) String() string {
	st := ""
	for y, li := range b {
		for x, v := range li {
			st += strconv.Itoa(v)
			st += " "
			if (x+1)%3 == 0 {
				st += " "
			}
		}
		if (y+1)%3 == 0 {
			st += "\n"
		}
		st += "\n"
	}
	return st
}
func (b Board) Eval(x, y, v int) bool {
	for i := 0; i < 9; i++ {
		// line
		if b.Get(i, y) == v {
			return false
		}
		// column
		if b.Get(x, i) == v {
			return false
		}
	}
	// square
	sqx := 3 * (x / 3)
	sqy := 3 * (y / 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Get(sqx+i, sqy+j) == v {
				return false
			}
		}
	}
	return true
}

func (g Board) Solve(sx, sy int) {
	for y := sy; y < 9; y++ {
		for x := sx; x < 9; x++ {
			if g.Get(x, y) != 0 {
				continue
			}
			for n := 1; n <= 9; n++ {
				if g.Eval(x, y, n) {
					g.Set(x, y, n)
					g.Solve(x+1, y)
					g.Set(x, y, 0)
				}
			}
			return
		}
		sx = 0
	}
	fmt.Println(g)
	fmt.Print("[enter]")
	fmt.Scanln()
}

func main() {
	g := &Board{}
	data := empty_json
	if len(os.Args) == 2 {
		name := os.Args[1]
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		data, err = io.ReadAll(f)
		if err != nil {
			panic(err)
		}
	}
	err := json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
	g.Solve(0, 0)
}
