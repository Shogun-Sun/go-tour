package themes

import "fmt"

func Interfaces() {
	g := 0.867 + 0.5i
	j := g
	fmt.Printf("Type: %T Value: %v\n", g, g)
	fmt.Printf("Type: %T\n", j)

	f := 3.142
	fmt.Printf("Type: %T Value: %v\n", f, f)

}
