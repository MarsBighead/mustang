package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0
var Lock sync.Mutex

func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {
		Lock.Lock()
		Counter = Counter + 1
		time.Sleep(1 * time.Nanosecond)
		Lock.Unlock()

	}

	Wait.Done()
}
