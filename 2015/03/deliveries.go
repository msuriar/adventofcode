package main

type Point struct{
	x, y int
}

type Course struct{
	position Point
	counts map[Point]int
}

func NewCourse() *Course {
	c := &Course{
		position: Point{x: 0, y: 0},
		counts: map[Point]int{},
	}
	c.counts[c.position] += 1

	return c
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
