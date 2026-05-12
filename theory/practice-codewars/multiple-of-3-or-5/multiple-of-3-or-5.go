package main

import "fmt"

func main() {
	Multiple3And5()
}

func Multiple3And5(number int) {
	var sum int
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Print(sum)
}
