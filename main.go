package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	const gs = 100
	//	counter := 0
	var counter int64
	//  var mu sync.Mutex
	var g sync.WaitGroup
	g.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			//  mu.Lock()
			atomic.AddInt64(&counter, 1)
			//	v := counter
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			//v++
			//	counter = v
			//  mu.Unlock()
			g.Done()

		}()
		fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())
	}
	g.Wait()
	fmt.Println("Num of CPu's: ", runtime.NumCPU())

	fmt.Println("count:", counter)

}
