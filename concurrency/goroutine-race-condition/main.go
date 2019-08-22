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
	for i := 0; i < max; i++ {
		wg.Add(1)
		go update(1, &wg)
	}

	wg.Wait()
	fmt.Println(x)
}

func update(i int, wg *sync.WaitGroup) {
	x += i
	wg.Done()
}

// END OMIT
