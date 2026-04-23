package themes

import (
	"fmt"
	"math"
)

func Condition() {
	fmt.Println(sqrt(-8))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}
