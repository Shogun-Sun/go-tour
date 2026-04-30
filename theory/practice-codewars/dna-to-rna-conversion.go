package main

import (
	"fmt"
)

func main() {
	var dna string = "GCAT"


	fmt.Println(DNAtoRNA(dna))
}

func DNAtoRNA(dna string) string {
	var rna string

	for _, j := range dna {
		if j == 71 {
			rna += "G"
		} else if j == 67 {
			rna += "C"
		} else if j == 65 {
			rna += "A"
		} else if j == 84 {
			rna += "U"
		}
	}

	return rna
}

