package themes

import (
	"fmt"
)

func Newton(z, x float64) float64 {
	return z - (z*z-x)/(2*z)
}

func Sqrt(x float64) float64 {
	z := 1.0
	iteration := 0
	// for i := 0; i < 10; i++ {
	// 	z = Newton(z, x)
	// 	fmt.Printf("Итерация %d: z = %g\n", i+1, z)
	// }

	for {
		prevZ := z
		z = Newton(z, x)
		diff := z - prevZ

		iteration++

		fmt.Printf("Итерация %d: z = %g\n", iteration, z)

		absoluteDiff := 0.0

		if diff < 0 {
			absoluteDiff = diff * -1.0
		} else {
			absoluteDiff = diff
		}

		if absoluteDiff/z < 1e-10 {
			break
		}

	}

	return z

}

func ExerciseLoopsAndFunctions() {
	fmt.Println(Sqrt(9104242152158291521))
}
