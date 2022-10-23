package main

import (
	"fmt"
	tty"github.com/mattn/go-tty"
)

func main() {
	t, _ := tty.Open()
	for {
		p, _ := t.ReadRune()
		if string(p) == "q" {
			break
		} else {
			fmt.Println(string(p))
		}
	}
}
