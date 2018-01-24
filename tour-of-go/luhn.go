package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v\n", luhn(8273123273520569))
	fmt.Printf("%v\n", luhn(4539148803436467))
	u := strings.Split("Hello", "")
	testSlices(u)
}

func luhn(creditNumber int) bool {
	tempNumber := creditNumber
	sum := 0

	for i := 1; i <= 16; i++ {
		digit := tempNumber % 10
		doubleDigit := 0

		if i%2 == 0 {
			doubleDigit = digit * 2
			if doubleDigit > 9 {
				doubleDigit -= 9
			}
			sum += doubleDigit
		} else {
			sum += digit
		}
		tempNumber = tempNumber / 10
	}

	return sum%10 == 0
}

func testSlices(sl []string) {
	for _, v := range sl {
		fmt.Println(v)
	}
}
