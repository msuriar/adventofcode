package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	w, d, h int
}

func NewPresent(s string) Present {
	p := Present{}

	dimensions := strings.Split(s, "x")

	numbers := [3]int{}

	for i, d := range dimensions {
		n, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		numbers[i] = n
	}

	p.w, p.d, p.h = numbers[0], numbers[1], numbers[2]
	return p
}

func (p *Present) FrontArea() int {
	return p.w * p.h
}

func (p *Present) SideArea() int {
	return p.h * p.d
}

func (p *Present) TopArea() int {
	return p.w * p.d
}

func (p *Present) Slack() int {
	n := p.FrontArea()

	if p.SideArea() < n {
		n = p.SideArea()
	}

	if p.TopArea() < n {
		n = p.TopArea()
	}
	return n
}

func paper(s string) int {
	p := NewPresent(s)
	return 2 * (p.FrontArea() + p.SideArea() + p.TopArea()) + p.Slack()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total_paper := 0

	for s.Scan() {
		total_paper += paper(s.Text())
	}

	fmt.Printf("Total wrapping paper: %d\n", total_paper)
}
