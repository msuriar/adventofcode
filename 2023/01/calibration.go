package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calibration_value(s string) int {
	nums := ""
	for _, rune := range s {
		if '0' <= rune && rune <= '9' {
			nums += string(rune)
		}
	}
	num := string(nums[0])+string(nums[len(nums)-1])
	r, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return r
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	total := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		total += calibration_value(s.Text())
	}
	fmt.Printf("Total calibration value: %d\n", total)
}
