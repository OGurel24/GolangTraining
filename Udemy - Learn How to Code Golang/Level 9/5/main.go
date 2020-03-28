package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64
	g := 200 // number of goroutines
	hold := 0
	wg.Add(g)

	for i := 0; i < g; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("End value:", hold)

}
