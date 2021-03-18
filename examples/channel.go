package examples

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ChannelTest() {
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	fmt.Println("==============")

	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("worker started!")

	fmt.Println("waiting for worker!")
	<-done
	fmt.Println("worker done")
}
