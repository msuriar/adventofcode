package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	lights [][]bool
}

type Point struct {
	x, y int
}

func NewGrid(x, y int) *Grid {
	l := make([][]bool, y)
	for i := range l {
		l[i] = make([]bool, x)
	}

	return &Grid{
		lights: l,
	}
}

func NewPoint(s string) Point {
	nums := strings.Split(s, ",")
	x, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	return Point{x, y}
}

func (g *Grid) Process(cmd string) {
	terms := strings.Split(cmd, " ")
	var start, end Point
	switch {
	case terms[0] == "turn":
		start = NewPoint(terms[2])
		end = NewPoint(terms[4])
	default:
		start = NewPoint(terms[1])
		end = NewPoint(terms[3])
	}

	switch {
	case terms[1] == "on":
		g.TurnOn(start, end)
	case terms[1] == "off":
		g.TurnOff(start, end)
	default:
		g.Toggle(start, end)
	}
}

func (g *Grid) TurnOn(start, end Point) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			g.lights[x][y] = true
		}
	}
}

func (g *Grid) TurnOff(start, end Point) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			g.lights[x][y] = false
		}
	}
}

func (g *Grid) Toggle(start, end Point) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			g.lights[x][y] = !g.lights[x][y]
		}
	}
}

func (g *Grid) NumOn() int {
	count := 0
	for _, x := range g.lights {
		for _, y := range x {
			if y {
				count++
			}
		}
	}
	return count
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	g := NewGrid(1000, 1000)

	for s.Scan() {
		g.Process(s.Text())
	}

	fmt.Printf("%v lights are on\n", g.NumOn())
}
