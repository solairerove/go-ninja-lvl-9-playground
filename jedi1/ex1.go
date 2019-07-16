package jedi1

import (
	"fmt"
	"sync"
)

// Ex1 tbd
func Ex1() {
	fmt.Println("\nEx1")

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println("hello from Tyler")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from Jack")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from Marla")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("\nexit")
}
