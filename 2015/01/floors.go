package main

func floor(s string) int {
	f := 0

	for _, c := range s {
		switch c {
		case '(':
		f++
		default:
		f--
		}
	}
	return f
}
