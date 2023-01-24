package random_examples_package

import (
	"fmt"
)

/*
An Example of Select in Golang:

It will print the first n fibonacci numbers
*/

func Fibonacci(n int) {
	c := make(chan int)
	quit := make(chan int)

	go func(c chan<- int, quit <-chan int) {
		x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("Quitting...")
				return
			}
		}

	}(c, quit)

	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
	close(quit) //Remember: closing always sends a message
}
