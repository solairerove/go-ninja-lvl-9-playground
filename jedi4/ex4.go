package jedi4

import (
	"fmt"
	"sync"
)

// Ex4 tbd
func Ex4() {
	fmt.Println("\nEx4")

	var wg sync.WaitGroup

	incrementor := 0
	gs := 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementor

			v++
			incrementor = v

			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(incrementor)
}
