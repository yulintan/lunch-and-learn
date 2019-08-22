package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// START OMIT
var x uint64
var max = 1000

func main() {
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		go update(1, &wg)
	}

	wg.Wait()

	fmt.Println(atomic.LoadUint64(&x))
}

func update(i int, wg *sync.WaitGroup) {
	atomic.AddUint64(&x, 1) // HLxxx
	wg.Done()
}

// END OMIT
