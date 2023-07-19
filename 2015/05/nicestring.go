package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsNice1(s string) bool {
	return EnoughVowels(s) && HasPair(s) && IsAllowed(s)
}

func IsNice2(s string) bool {
	return HasDuplicatePair(s) && Has3Palindrome(s)
}

func HasDuplicatePair(s string) bool {
	for i := 0 ; i < len(s) - 1; i++ {
		p := s[i:i+2]
		if strings.Count(s, p) > 1 {
			return true
		}
	}
	return false
}

func Has3Palindrome(s string) bool {
	for i := 0 ; i < len(s) - 2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func EnoughVowels(s string) bool {
	vowels := []string{
		"a",
		"e",
		"i",
		"o",
		"u",
	}
	count := 0

	for _, v := range vowels {
		count += strings.Count(s, v)
	}

	return count >= 3
}

func HasPair(s string) bool {
	for i := 0; i < len(s)-1 ; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func IsAllowed(s string) bool {
	forbidden := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}

	for _, f := range forbidden {
		if strings.Contains(s, f) {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	nice := 0
	nice2 := 0

	for s.Scan() {
		if IsNice1(s.Text()) {
			nice++
		}
		if IsNice2(s.Text()) {
			nice2++
		}
	}

	fmt.Printf("Input contained %d nice strings.\n", nice)
	fmt.Printf("Input contained %d nice2 strings.\n", nice2)
}
