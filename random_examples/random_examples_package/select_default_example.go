package random_examples_package

import (
	"fmt"
	"time"
)

func TickTickBoom() {
	tick := time.NewTicker(time.Second * 1) //time.Tick() is a wrapper over NewTicker, but it leaks
	boom := time.After(time.Second * 5)

	for {
		select {
		case <-tick.C:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			time.Sleep(time.Millisecond * 200)
			fmt.Println(".")
		}
	}
}
