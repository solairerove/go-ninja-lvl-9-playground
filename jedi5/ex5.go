package jedi5

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Ex5 tbd
func Ex5() {
	fmt.Println("\nEx5")

	var wg sync.WaitGroup

	var incrementor int64
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(incrementor)
}
