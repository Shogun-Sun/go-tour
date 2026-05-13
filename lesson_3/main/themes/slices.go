package themes

import "fmt"

func Slices() {
	primes := [6]int{2, 3, 5, 7, 1, 13}

	var s []int = primes[1:3]
	fmt.Println(s)

	s = primes[:5]
	fmt.Println(s)

	s = primes[1:]
	fmt.Println(s)

	s = primes[:]
	fmt.Println(s)
}
