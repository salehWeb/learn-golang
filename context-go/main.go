package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"unsafe"
)

func main() {
	fmt.Printf("context allow you to cancel the operation after certain time\n\n")

	timeout, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	req, err := http.NewRequestWithContext(timeout, "GET", "https://github.com", nil)

	if err != nil {
		panic(err) // return  Get "https://github.com": context deadline exceeded
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Printf("\nResponse Size is %d bytes\n", unsafe.Sizeof(*res))
	fmt.Println("Server is :", res.Header["Server"][0])
}
