package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	g := 200 // number of goroutines
	hold := 0
	wg.Add(g)

	for i := 0; i < g; i++ {
		go func() {
			m.Lock()
			fmt.Println(hold)
			v := hold
			v++
			hold = v
			wg.Done()
			m.Unlock()
		}()

	}

	wg.Wait()
	fmt.Println("End value:", hold)

}