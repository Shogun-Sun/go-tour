package themes

import "fmt"

const Pi = 3.14

func Constants() {
	// Pi = 20 // error

	const truth = true

	fmt.Println(Pi, truth)
}
