package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

const N = 3000000
const n = N / 10

func howMany(i int) {
	if i%n == 0 {
		// fmt.Printf("I've handled %d requests\n", i)
	}
}

//START OOM2 OMIT
func consumer(c <-chan struct{}) {
	for i := 1; i <= N; i++ {
		<-c
		howMany(i)
	}
}

func producer(c chan<- struct{}) {
	c <- struct{}{}
}

//END OOM2 OMIT

var theChan = make(chan struct{})

func prof() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func report() {
	var m runtime.MemStats
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("There are %d goroutines alive\n", runtime.NumGoroutine())

			runtime.ReadMemStats(&m)
			fmt.Printf("Alloc: %d\n totalalloc: %d\n", m.Alloc, m.TotalAlloc)
			fmt.Printf("sys: %d\n", m.Sys)
			fmt.Printf("stackinuse: %d\n", m.StackInuse)
			fmt.Printf("stacksys: %d\n", m.StackSys)
		}
	}()
}

func main() {
	prof()
	report()

	//START OOM OMIT
	for i := 0; i < N; i++ {
		//N is 3,000,000 // HL
		go producer(theChan)
	}

	time.Sleep(5 * time.Second)

	consumer(theChan)
	//END OOM OMIT

	time.Sleep(1 * time.Hour)

}
