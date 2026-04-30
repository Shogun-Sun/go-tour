package main

import (
	"fmt"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David"}

	// Range возвращает две переменные:
	// 1 - индекс элемента в срезе
	// 2 - значение элемента

	for i, name := range names {
		fmt.Printf("id: %d, name: %s\n", i, name)
	}
}