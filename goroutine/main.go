package main

import (
	"fmt"
	"net/http"
	"time"
)

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
		go SendHttpRequest(webSites[i])
	}

	time.Sleep(time.Second)
	fmt.Println(time.Now())
}

func SendHttpRequest(url string) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("status code is: %d and content type is %s\n\n", res.StatusCode, res.Header.Values("Content-Type"))
	defer res.Body.Close()
}
