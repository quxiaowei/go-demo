package examples

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ChannelRange() {
	queue := make(chan string, 5)

	go func() {
		t := time.NewTicker(1 * time.Second)
		for range t.C {
			fmt.Println(".")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			queue <- strconv.Itoa((i + 1))
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		}
		close(queue)
		fmt.Println("channel closed!")
	}()

	for s := range queue {
		fmt.Println(s)
	}

	fmt.Println("finished!")
}
