package random_examples_package

import (
	"fmt"
	"time"
)

/*
Channels can be considered as data streams.
If that is the case, then we can use select to manipulate data streams.

This code is an example of Fan out. It works like a load balancer.
*/

func fanOut(in <-chan int, outA, outB chan<- int){
	for data:= range in {
		select {
		case outA <- data:
		case outB <- data:
		}
	}
}

func worker(out <-chan int, workerName string){
	for {
		fmt.Printf("Worker %s working: num received %d\n", workerName, <-out)
		time.Sleep(time.Second * 2)
	}
}

func FanoutExample(){
	in := make(chan int)
	defer close(in)
	outA := make(chan int)
	outB := make(chan int)

	go worker(outA, "A")
	go worker(outB, "B")

	go fanOut(in, outA, outB)

	for i:=0; i<10; i++{
		in<-i
	}
	time.Sleep(time.Second * 2)
}