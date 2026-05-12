package themes

import "fmt"

func StackingDefers() {
	fmt.Println("counting")
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}

}
