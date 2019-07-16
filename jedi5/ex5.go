package jedi5

import (
	"fmt"
	"runtime"
	"sync"
)

// Ex5 tbd
func Ex5() {
	fmt.Println("\nEx5")

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
