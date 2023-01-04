package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func check_nil_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Loading...")

	var user_name string = "salehWeb"
	var url string = "https://api.github.com/users/" + user_name + "/repos?page=1&per_page=1"
	var file_name = "./" + user_name + ".json"

	res, err := http.Get(url)
	check_nil_err(err)

	bytes, err := io.ReadAll(res.Body)

	check_nil_err(err)

	var json string = string(bytes)

	file, err := os.Create(file_name)
	check_nil_err(err)

	io.WriteString(file, json)

	fmt.Println("Done.")

	defer file.Close()
	defer res.Body.Close()
}
