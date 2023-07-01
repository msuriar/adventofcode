package main

import (
	"bufio"
	"fmt"
	"os"
)

func floor(s string) int {
	f := 0

	for _, c := range s {
		switch c {
		case '(':
			f++
		default:
			f--
		}
	}
	return f
}

func basement(s string) int {
	f := 0

	for i, c := range s {
		switch c {
		case '(':
			f++
		default:
			f--
		}
		if f < 0 {
			return i+1
		}
	}
	return f
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	if s.Scan() {
		dst := floor(s.Text())
		fmt.Printf("Santa ended at floor %d\n", dst)
	}
}
