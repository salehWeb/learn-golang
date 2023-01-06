package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var num uint32 = 0
	mutex := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				num++
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Num End Value Is: %d", num) // it should return 1_000_000
	// without mutex it return number less the 1_000_000

	var SampleMap map[int]int = map[int]int{}

	// go writeLoopError(SampleMap)
	// go readLoopError(SampleMap)

	// this code above return Error "fatal error: concurrent map iteration and map write"

	// this is The Good Code in this case
	RWMutex := &sync.RWMutex{}

	go writeLoopError(SampleMap, RWMutex)
	go readLoopError(SampleMap, RWMutex)

	time.Sleep(5 * time.Second)
}

func writeLoopError(m map[int]int, RWMutex *sync.RWMutex) {
	for {
		RWMutex.Lock()
		for i := 0; i < 100; i++ {
			m[i] = i
		}
		RWMutex.Unlock()
	}
}

func readLoopError(m map[int]int, RWMutex *sync.RWMutex) {
	for {
		RWMutex.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		RWMutex.RUnlock()
	}
}
