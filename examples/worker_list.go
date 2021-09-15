package examples

import (
	"math/rand"
)

type Worker struct {
	name  string
	jobId int
}

func GetWorkers(count int) []Worker {

	name_list := []string{"Jone", "Jane", "Mike",
		"Smith", "Lee", "Porter", "Shawn", "Seven", "Eason",
		"Ben", "Brown", "Bob", "Nick", "Petty", "Max", "Kate", "Eric",
		"Carl", "Dick", "Eddie", "Harry", "Ivan", "Julie", "Lisa",
		"Mary", "Neil", "Sandy", "Susan", "Tom", "Tim", "Simon",
		"Anna", "Coco", "Cindy", "Davie", "Frank", "Lauras",
		"Lucy", "Nancy", "May", "Paul", "Sam", "Shirley", "Sean"}

	result := make([]Worker, 0)
	perm := rand.Perm(len(name_list))
	for _, i := range perm[:count] {
		result = append(result, Worker{name: name_list[i]})
	}
	return result
}
