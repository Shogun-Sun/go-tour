package themes

import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Printf("Значение i: %d\n", *p)

	*p = 21
	fmt.Printf("Новое значение i: %d\n", *p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
