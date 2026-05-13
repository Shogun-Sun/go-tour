package themes

import "fmt"

func CreatingASliceWithMake() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s\nlen: %d\ncap: %d\n%v\n", s, len(x), cap(x), x)
}
