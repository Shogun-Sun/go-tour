package main

import "fmt"


func main() {
	fmt.Println(StringToArray("I Love Go"))
}

func StringToArray(str string) []string {
	var result []string
	var temp_word string

	
	for i, e := range str {
		switch string(e) {
			case " ": 
				result = append(result, temp_word)
				temp_word = ""
			default:
				temp_word += string(e) 
		}

		if i == len(str) - 1 {
			result = append(result, temp_word)
			temp_word = ""
		}
	}

	fmt.Println(result)

	return result
}