package main

import "fmt"

func checkOnLamps() int {
	data := make(map[int]bool)
	var countingLamps = 0

	for init := 1; init <= 100; init++ {
		data[init] = true
	}

	for i := 1; i <= 100; i++ {
		if i == 1 {
			continue
		} else {
			for k := i; k <= 100; k *= i {
				if data[k] == true {
					data[k] = false
				} else {
					data[k] = true
				}
			}
		}
	}

	for _, value := range data {
		if value == true {
			countingLamps += 1
		}
	}

	return countingLamps
}

func main() {
	fmt.Println(checkOnLamps())
}
