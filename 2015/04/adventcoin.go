package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func FindHashSuffix(s string, p int) int {
	slice := []string{}
	for i := 0 ; i < p ; i++ {
		slice = append(slice, "0")
	}
	prefix := strings.Join(slice, "")

	h := md5.New()
	suffix := 1
	for {
		candidate := fmt.Sprintf("%s%d", s, suffix)
		io.WriteString(h, candidate)
		hash := fmt.Sprintf("%x", h.Sum(nil))

		if strings.HasPrefix(hash, prefix) {
			return suffix
		}
		h.Reset()
		suffix++
	}
}

func main() {
	input := "bgvyzdsv"

	five := FindHashSuffix(input, 5)
	six := FindHashSuffix(input, 6)

	fmt.Printf("FindHashSuffix(%q, 5) => %d\n", input, five)
	fmt.Printf("FindHashSuffix(%q, 6) => %d\n", input, six)
}
