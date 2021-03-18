package examples

import (
	"fmt"
	"strconv"
	"time"
)

func ChannelRange() {
	queue := make(chan string, 5)

	go func() {
		t := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-t.C:
				fmt.Println(".")
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			queue <- strconv.Itoa((i + 1))
		}

		<-time.After(10 * time.Second)
		close(queue)
		fmt.Println("channel closed!")
	}()

	for s := range queue {
		fmt.Println(s)
	}

	fmt.Println("finished!")
}
