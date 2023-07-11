package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Course struct {
	position Point
	counts   map[Point]int
}

func NewCourse() *Course {
	c := &Course{
		position: Point{x: 0, y: 0},
		counts:   map[Point]int{},
	}
	c.counts[c.position] += 1

	return c
}

func (c *Course) PartnerNavigate(tc string) {
	santa := []string{}
	robosanta := []string{}

	for i, r := range tc {
		if i%2 == 0 {
			santa = append(santa, string(r))
		} else {
			robosanta = append(robosanta, string(r))
		}
	}

	c.Navigate(strings.Join(santa, ""))
	c.position = Point{x: 0, y: 0}
	c.Navigate(strings.Join(robosanta, ""))
}

func (c *Course) Navigate(tc string) {
	for _, r := range tc {
		switch r {
		case '>':
			c.position.x += 1
		case 'v':
			c.position.y -= 1
		case '<':
			c.position.x -= 1
		case '^':
			c.position.y += 1
		}

		c.counts[c.position] += 1
	}

}

func (c *Course) UniqueHouses() int {
	return len(c.counts)
}

func main() {
	c := NewCourse()

	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	c.Navigate(string(b))

	fmt.Printf("Single worker, unique houses: %d\n", c.UniqueHouses())

	c2 := NewCourse()
	c2.PartnerNavigate(string(b))

	fmt.Printf("Two workers, unique houses: %d\n", c2.UniqueHouses())
}
