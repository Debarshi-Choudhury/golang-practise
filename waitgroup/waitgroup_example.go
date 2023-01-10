package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Finished goroutine")
	wg.Done()
}

func main() {
	fmt.Println("Go WaitGroup example:")

	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait() //blocks until 0
	fmt.Println("Finished Program")
}
