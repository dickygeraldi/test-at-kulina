package main

import "fmt"

func sockMerchant(n int, ar []int) int {
	dataSock := make(map[int]int)
	var countingSock = 0

	if n == len(ar) {
		for i := 0; i < int(len(ar)); i++ {
			if dataSock[int(ar[i])] == 0 {
				dataSock[int(ar[i])] = 1
			} else {
				dataSock[int(ar[i])] += 1
			}
		}
	} else {
		return -1
	}

	for _, value := range dataSock {
		if value >= 2 {
			countingSock += value / 2
		}
	}
	return countingSock
}

func main() {
	data := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(len(data), data))
}
