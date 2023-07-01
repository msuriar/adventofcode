package main

import (
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
	return p.w*p.h
}

func (p *Present) SideArea() int {
	return p.h*p.d
}

func (p *Present) TopArea() int {
	return p.w*p.d
}

func paper(s string) int {
	p := NewPresent(s)
	return 2 * (p.FrontArea() + p.SideArea() + p.TopArea())
}
