package main

import "fmt"
var literPerHour = 0.5


func main() {
	Litres(11.8)
}


func Litres(time float64) {
	fmt.Println(int(literPerHour * time))
}