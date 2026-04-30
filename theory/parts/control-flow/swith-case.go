package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Nika"}

	for _, name := range names {
		switch name {
			case "Alice": 
				fmt.Println("girl")
			case "Bob": 
				fmt.Println("boy")
			case "Charlie":
				fmt.Println("boy")
			case "Nika": 
				fmt.Println("girl")
			default:
				fmt.Println("unknown")
		}
	}
}