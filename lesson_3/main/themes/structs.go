package themes

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Structures() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = v.Y + v.X
	v.Y = v.X + v.Y
	fmt.Printf("%d %d\n", v.X, v.Y)
}
