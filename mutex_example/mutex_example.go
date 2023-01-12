package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func deposit(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Adding %d to balance %d\n", amount, balance)
	balance += amount
	mutex.Unlock()
	wg.Done()
}

func withdraw(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Removing %d from balance %d\n", amount, balance)
	balance -= amount
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Mutex example:")
	//initial balance
	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(500, &wg)
	go withdraw(700, &wg)
	wg.Wait()

	fmt.Printf("Final balance is %d\n", balance)
}
