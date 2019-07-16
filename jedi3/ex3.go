package jedi3

import (
	"fmt"
	"runtime"
	"sync"
)

// Ex3 tbd
func Ex3() {
	fmt.Println("\nEx3")

	var wg sync.WaitGroup

	incrementor := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementor

			runtime.Gosched()

			incrementor = v

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(incrementor)
}
