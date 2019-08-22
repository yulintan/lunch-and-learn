package main

import (
	"fmt"
	"runtime"
	"time"
)

// START OMIT
var max = 10

func main() {
	for i := 0; i < max; i++ {
		go myfunc(i)
	}

	time.Sleep(1 * time.Second)
}

func myfunc(id int) {
	fmt.Printf("id = %d\n", id)
	for {
		runtime.Gosched() // HLxxx
	}
}

// END OMIT
