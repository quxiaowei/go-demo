package examples

import (
	"fmt"
	"time"
)

func Switch() {
	switch i := 1; {
	case i == 1:
		fmt.Println("one")
	case i == 2:
		fmt.Println("two")
	case i == 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
