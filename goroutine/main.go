package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// // goroutine
	// // goroutine is Just Awesome Man 🤩

	var webSites [3]string = [3]string{
		"https://pkg.go.dev/net/http",
		"https://github.com",
		"https://go.dev/doc/effective_go",
	}

	fmt.Println(time.Now())

	for i := 0; i < len(webSites); i++ {
		wg.Add(1)
		go SendHttpRequest(webSites[i])
	}

	wg.Wait()

	fmt.Println(time.Now())
}

func SendHttpRequest(url string) {
	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("status code is: %d and content type is %s\n\n", res.StatusCode, res.Header.Values("Content-Type"))
	defer res.Body.Close()
}
