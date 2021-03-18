package examples

import "fmt"

func Structs() {
	type person struct {
		name string
		age  int
	}

	newPerson := func(name string) *person {
		p := person{name: name}
		p.age = 42
		return &p
	}

	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(&person{name: "Alice", age: 30})
	fmt.Println(newPerson("john"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)
}
