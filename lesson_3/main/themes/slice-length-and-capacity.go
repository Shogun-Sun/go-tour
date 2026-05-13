package themes

import "fmt"

func SliceLengthAndCapacity() {
	var s = []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%sИсходный срез:%s\n", "\033[31m", "\033[0m")
	printSlice(s)

	fmt.Printf("%sПривести к нулевому размеру:%s\n", "\033[31m", "\033[0m")
	s = s[:0]
	printSlice(s)

	fmt.Printf("%sУвеличить длину:%s\n", "\033[31m", "\033[0m")
	s = s[:4]
	printSlice(s)

	fmt.Printf("%sОтбросить первые два значения:%s\n", "\033[31m", "\033[0m")
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len: %d\ncap: %d\n%v\n", len(s), cap(s), s)
}
