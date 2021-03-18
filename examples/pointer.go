package examples

import "fmt"

func Pointer() {
	zeroval := func(ival int) {
		ival = 0
	}

	zeroptr := func(iptr *int) {
		*iptr = 0
	}

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", i)
}
