package themes

import "fmt"

func AppendingToASlice() {
	var s []int
	printSlice3(s)

	s = append(s, 0)
	printSlice3(s)

	s = append(s, 1)
	printSlice3(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice3(s []int) {
	fmt.Printf("len: %d\ncap: %d\n%v\n", len(s), cap(s), s)
}
