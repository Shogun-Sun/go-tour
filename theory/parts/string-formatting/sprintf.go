package main 

import "fmt"

func main() {
	fmt.Println(Greet("Alice"))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}