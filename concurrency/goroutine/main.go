package main

import (
	"fmt"
	"time"
)

// START OMIT
var max = 1000

func main() {
	c := make(chan int) // HLxxx
	for i := 0; i < max; i++ {
		go update(1, c)
	}

	c <- 0 // HLxxx
	time.Sleep(time.Millisecond)
	x := <-c // HLxxx
	close(c)
	fmt.Println(x)
}

func update(i int, c chan int) {
	sum := <-c // HLxxx
	sum += i
	c <- sum // HLxxx
}

// END OMIT
