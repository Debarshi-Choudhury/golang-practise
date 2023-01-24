package random_examples_package

import (
	"fmt"
	"time"
)

/*
Making channels non-blocking using select
*/

func tryReceive(c <-chan int) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	default: //processes when c is blocking
		return 0, true, false
	}
}

func TryReceiveExample() {
	c := make(chan int)
	go func(c chan<- int) {
		for {
			c <- 1
			time.Sleep(time.Second * 1)
		}
	}(c)

	for i := 0; i < 20; i++ {
		data, more, ok := tryReceive(c)
		fmt.Println(data, more, ok)
		time.Sleep(time.Millisecond * 200)
	}
}
