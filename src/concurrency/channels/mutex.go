package channels

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}

func Mutex() {
	var wg sync.WaitGroup
	var m sync.Mutex // If mutex is not used, there would be race condition and output would be less than 1000

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	} 

	wg.Wait()

	fmt.Println("Value of i after 1000 operations is ", i)
}