package themes

import (
	"fmt"
)

var a, b int = 10, 20

// Variables with initializers
var d, e = false, "test"

func Variables() {
	var c int = 30

	//Short variable declarations
	f := false

	fmt.Println(a, b, c, d, e, f)
}
