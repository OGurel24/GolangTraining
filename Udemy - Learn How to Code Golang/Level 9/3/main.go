package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	g := 200 // number of goroutines
	hold := 0
	wg.Add(g)

	for i := 0; i < g; i++ {
		go func() {
			fmt.Println(hold)
			v := hold
			runtime.Gosched()
			v++
			hold = v
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("End value:", hold)

}
