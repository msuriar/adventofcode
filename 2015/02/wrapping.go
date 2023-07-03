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

func (p *Present) Volume() int {
	return p.w * p.h * p.d
}

func (p *Present) SmallPerimeter() int {
	s1 := 2*(p.w + p.h)
	s2 := 2*(p.w + p.d)
	s3 := 2*(p.h + p.d)

	if s1 <= s2 && s1 <= s3 {
		return s1
	} else if s2 <= s1 && s2 <= s3 {
		return s2
	}
	return s3
}

func paper(s string) int {
	p := NewPresent(s)
	return 2 * (p.FrontArea() + p.SideArea() + p.TopArea()) + p.Slack()
}

func ribbon(s string) int {
	p := NewPresent(s)

	return p.Volume() + p.SmallPerimeter()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total_paper := 0
	total_ribbon := 0

	for s.Scan() {
		total_paper += paper(s.Text())
		total_ribbon += ribbon(s.Text())
	}

	fmt.Printf("Total wrapping paper: %d\n", total_paper)
	fmt.Printf("Total ribbon: %d\n", total_ribbon)
}
