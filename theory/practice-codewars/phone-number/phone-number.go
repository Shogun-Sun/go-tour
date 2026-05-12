package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	CreatePhoneNumber([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
}

func CreatePhoneNumberWithBuilder(nums [10]int) string {
	var b strings.Builder
	b.Grow(14)

	b.WriteByte('(')
	for i := 0; i < 3; i++ {
		b.WriteString(strconv.Itoa(int(nums[i])))
	}

	b.WriteString(") ")
	for i := 3; i < 6; i++ {
		b.WriteString(strconv.Itoa(int(nums[i])))
	}

	b.WriteByte('-')
	for i := 6; i < 10; i++ {
		b.WriteString(strconv.Itoa(int(nums[i])))
	}

	return b.String()
}

func CreatePhoneNumber(nums [10]int) {
	var finalString string
	for o, i := range nums {
		switch o {
		case 0:
			finalString += "("
		case 3:
			finalString += ") "
		case 6:
			finalString += "-"
		}

		finalString += strconv.Itoa(i)
	}
	fmt.Println(finalString)
}
