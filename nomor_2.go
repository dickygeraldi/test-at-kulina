package main

import "fmt"

func countingValleys(n int, s string) int {
	var countingValley = 0
	var now = 0
	var before = 0
	getLength := len(s)

	if getLength == n {

		for i := 0; i < int(n); i++ {
			before = now

			if string(s[i]) == "U" {
				now += 1
				if now == 0 && before < 0 {
					countingValley += 1
				}
			} else if string(s[i]) == "D" {
				now -= 1
			}

		}
	} else {
		return -1
	}

	return countingValley
}

func main() {
	data := "UDDDUDUU"
	fmt.Println(countingValleys(len(data), data))
}
