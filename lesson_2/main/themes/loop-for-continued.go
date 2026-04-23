package themes

import "fmt"

func LoopForContinued() {
	sum := 1
	for sum < 900 {
		sum += sum
	}

	fmt.Println(sum)
}
