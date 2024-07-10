package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(IsValid("1234567890121"))
}

func IsValid(id string) bool {
	list := []rune(id)
	sum := 0
	var validNumber int
	for i, v := range list {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		if i != 12 {
			sum += (13 - i) * num //multiply number with its reversed index
		} else {
			validNumber = num //last digit
		}
	}
	return ((11-(sum%11))%10 == validNumber)
}
