package examples

import "fmt"

func Range() {
	nums := []int{2, 3, 4}
	sum := 0

	for i, num := range nums {
		sum += num
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
