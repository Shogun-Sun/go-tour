package main

import "fmt"

func main() {
	fmt.Println(Solution("world"))
}

func Solution (str string) string {
	var result string
	var temp_array []string

	for _, e := range str {
		temp_array = append(temp_array, string(e))
	}

	for i := len(temp_array) - 1; i >= 0; i-- {
		result += temp_array[i]
	} 

	return result
}