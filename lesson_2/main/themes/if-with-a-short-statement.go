package themes

import "fmt"

func IfShortStatement() {
	if v := 10; v < 20 {
		fmt.Println("v меньше 20 и равно:", v)
	}
}
