package examples

import "fmt"

func Sum() {
	_Sum := func(nums ...int) {
		fmt.Print(nums, "")
		total := 0
		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}

	_Sum(1, 2)
	_Sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	_Sum(nums...)
}
