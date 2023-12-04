package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func calibration_value(s string) int {
	digit_re := "(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)"
	re := regexp.MustCompile(digit_re)

	matches := re.FindAllString(s, -1)

	first, last := text_to_digit(matches[0]), text_to_digit(matches[len(matches)-1])

	num := first+last

	r, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	return r
}

func text_to_digit(s string) string {
	text_re := "one|two|three|four|five|six|seven|eight|nine"

	digit := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	re := regexp.MustCompile(text_re)

	if re.MatchString(s) {
		return digit[s]
	}
	return s
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
