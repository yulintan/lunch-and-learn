package main

import (
	"fmt"
	"sync"
)

// START OMIT
var x = 0
var max = 1000

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex
	for i := 0; i < max; i++ {
		wg.Add(1)
		go update(1, &wg, &mux)
	}

	wg.Wait()
	fmt.Println(x)
}

func update(i int, wg *sync.WaitGroup, mux *sync.Mutex) {
	mux.Lock() // HLxxx
	x += i
	mux.Unlock() // HLxxx

	wg.Done()
}

// END OMIT
