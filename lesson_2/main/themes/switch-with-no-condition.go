package themes

import (
	"fmt"
	"time"
)

func SwitchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Доброе утро!")
	case t.Hour() < 17:
		fmt.Printf("%d:%d\n", t.Hour(), t.Minute())
		fmt.Println("Добрый день.")
	default:
		fmt.Println("Добрвый вечер.")
	}
}
