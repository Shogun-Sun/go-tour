package main

import "fmt"

func main() {
	var x = 10
	var y = 20

	swap(&x, &y)

	fmt.Println(x, y)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
