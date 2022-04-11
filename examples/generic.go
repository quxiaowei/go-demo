package examples

import "fmt"

type number interface{
	int64 | float64
}

func sum[K comparable, V int64 | float64] (m map[K]V) V{
	var s V
	for _, v := range m{
		s += v
	}
	return s
}

func sum2[K comparable, V number] (m map[K]V) V{
	var s V
	for _, v := range m{
		s += v
	}
	return s
}

func Generic(){
	ints := map[string]int64 {
		"first":	34,
		"second":	12,
	}
	floats:= map[string]float64 {
		"first": 35.32,
		"second": 25.45,
	}

	fmt.Printf("int sum: %v", sum(ints))
	fmt.Println()
	fmt.Printf("float sum: %v", sum(floats))
}