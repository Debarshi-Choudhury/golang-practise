package random_examples_package

/*
An example to see that once you close a channel, no matter how many times
you read from it, it will give you the zero value of the channel type and a bool "false" which
indicates that the channel has been closed.
*/

import (
	"fmt"
	"time"
)

func myFunc(c <-chan int) {
	fmt.Println("myFunc Starts")
	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
	fmt.Println("myFunc Ends")
}

func ReadingFromChannelAfterBeingClosed() {
	fmt.Println("Main Starts:")
	c := make(chan int)
	go myFunc(c)
	c <- 99
	close(c)
	time.Sleep(5 * time.Second)

	fmt.Println("Main Ends")
}
