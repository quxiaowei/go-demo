package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func ChannelSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		c2 <- "Two"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout! Channel One")
	}

	select {
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout! Channel two")
	}

}
