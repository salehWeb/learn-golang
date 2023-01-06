package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {

	// unBuffered Channel
	chan1 := make(chan string)

	go func() {
		chan1 <- "Hello World"
	}()

	msg1 := <-chan1

	fmt.Printf("Massage From unBuffered Channel Is: \"%s\"\n\n", msg1)

	close(chan1)

	// Buffered Channel

	chan2 := make(chan string, 1)

	chan2 <- "Hello World"

	msg2 := <-chan2

	fmt.Printf("Massage From Buffered Channel Is: \"%s\"\n\n", msg2)

	close(chan2)

	chan3 := make(chan string)

	go func() {
		wg := sync.WaitGroup{}
		for i := 1; i < 1001; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				msg := GetMassage()
				chan3 <- msg
			}()
		}
		wg.Wait()
		close(chan3)
	}()

	for msg := range chan3 {
		fmt.Printf("Massage From Channel 3 Is: \"%s\"\n\n", msg)
	}
}

func GetMassage() string {
	time.Sleep(time.Second)
	rand.Seed(time.Now().UnixMicro())
	return "Hello World" + strconv.FormatInt(rand.Int63(), 10)
}
