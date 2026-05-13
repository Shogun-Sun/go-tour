package themes

import "fmt"

func SliceLiterals() {
	// var q = []int{2, 3, 5, 7, 11, 13}
	var r = []bool{true, false, true, true, false, true}

	r = append(r, false)

	fmt.Println(r)

	var s = []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	s = append(s, struct {
		i int
		b bool
	}{15, false})
	fmt.Println(s)
}
