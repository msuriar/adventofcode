package main

import "strconv"

func calibration_value(s string) int {
	nums := ""
	for _, rune := range s {
		if '0' <= rune && rune <= '9' {
			nums += string(rune)
		}
	}
	num := string(nums[0])+string(nums[len(nums)-1])
	r, _ := strconv.Atoi(num)
	return r
}
