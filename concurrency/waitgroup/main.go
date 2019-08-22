package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
var x = 0
var max = 1000

func main() {
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		go update(1, &wg)
		wg.Wait()
	}

	time.Sleep(time.Second)
	fmt.Println(x)
}

func update(i int, wg *sync.WaitGroup) {
	x += i

	wg.Done()
}

// END OMIT
