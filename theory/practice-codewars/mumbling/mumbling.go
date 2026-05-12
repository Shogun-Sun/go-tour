package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Print(Accum("aBcD"))
}

func Accum(s string) string {
	var resultString string
	var formatString string
	for _, e := range s {
		formatString += string(unicode.ToLower(e))
	}

	for i, e := range formatString {
		for j := 0; j <= i; j++ {
			if j == 0 {
				resultString += string(unicode.ToUpper(e))
			} else {
				resultString += string(e)
			}
		}

		if i < len(s)-1 {
			resultString += "-"
		}
	}

	return resultString
}
