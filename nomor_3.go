package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func separatingNumber(number string) {
	reg, _ := regexp.Compile("[^0-9]")
	stringProccess := reg.ReplaceAllString(number, "")
	var pencacah = len(stringProccess)

	if pencacah < 19 {
		satuan := [19]int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000, 1000000000000000, 10000000000000000, 100000000000000000, 1000000000000000000}

		bilangan, _ := strconv.Atoi(stringProccess)
		var newbil = satuan[pencacah-1]
		var init = 0

		for {
			init = bilangan / newbil
			fmt.Println(newbil * init)
			bilangan -= init * newbil
			newbil /= 10

			if newbil == 0 {
				break
			}
		}
	} else {
		fmt.Println("Digit bilangan harus kurang dari 19 digit")
	}
}

func main() {
	numberData := "1.345.679"
	separatingNumber(numberData)
}
