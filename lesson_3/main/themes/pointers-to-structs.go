package themes

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func PointersToStructs() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9

	fmt.Println(v)
}
