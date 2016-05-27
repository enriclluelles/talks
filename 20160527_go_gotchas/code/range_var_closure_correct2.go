package main

import (
	"fmt"
	"time"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// START OMIT
	for _, e := range sl {
		a := e
		go func() {
			fmt.Printf("Hello from iteration %d\n", a)
		}()
	}
	// END OMIT

	time.Sleep(1 * time.Second)
}
