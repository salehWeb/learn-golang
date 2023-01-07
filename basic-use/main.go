package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type user struct {
	name    string
	isAdmin bool
	id      int32
}

func defer_function() {
	defer fmt.Printf("World-3\n")
}

func is_authorized(user user) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("we had been recovered...\n\n")
		}
	}()

	if !user.isAdmin {
		panic(errors.New("unAuthorized to this action"))
		fmt.Printf("It Will Not Print\n\n")
	} else {
		fmt.Printf("%s You Good sir", user.name)
	}
}

func main() {
	fmt.Printf("\n\nFor Loop\n\n")

	for i := 0; i < 10; i++ {
		fmt.Printf("\ni: %d", i)
	}

	fmt.Printf("\n\nRange \n\n")
	fmt.Printf("Range is a form of for loop that can be used with slices or maps\n\n")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2*%d = %d\n", i, v)
	}

	fmt.Printf("\nif Statement\n\n")

	var a int8 = 16
	var b int8 = 15

	for i := 0; i < 3; i++ {
		fmt.Printf("a is %d\n", a)
		fmt.Printf("b is %d\n", b)
		if a > b {
			fmt.Printf("a bigger then b\n")
		} else if a < b {
			fmt.Printf("a smaller then b\n")
		} else {
			fmt.Printf("a equal to b\n")
		}

		b++
		fmt.Printf("\n\n")
	}

	fmt.Printf("\n Switch Case \n\n")

	var c string = "a"

	for i := 0; i < 5; i++ {
		fmt.Printf("c is %s \n\n", c)
		switch c {
		case "a":
			c = "b"
		case "b":
			c = "c"
		case "c":
			c = "d"
		default:
			c = "End"
		}
	}

	fmt.Printf("\n Error handling \n\n")
	src, err := os.ReadFile("./basic-use/text.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("file content is \"%s\"", strings.TrimSpace(string(src))) // after reading file you need converted it from byte array to string
	}

	fmt.Printf("\n\nPanic and Recover \n\n")
	fmt.Printf("Panic will disable any function that bellow it when it called\n\n")
	fmt.Printf("Panic will not disable deferred function that in top of it\n\n")
	fmt.Printf("Panic will run deferred function after panic massage print\n\n")

	var user1 user = user{
		name:    "wefew effwe",
		id:      1,
		isAdmin: true,
	}

	var user2 user = user{
		name:    "oihiio oioio",
		id:      2,
		isAdmin: false,
	}

	var user3 user = user{
		name:    "iowu asoihv",
		id:      1,
		isAdmin: true,
	}

	is_authorized(user1)
	is_authorized(user2)
	is_authorized(user3)

	fmt.Printf("\n\nBuffer \n\n")
	fmt.Printf("A Buffer is a variable-sized buffer of bytes with Read and Write methods.\n\n")

	var buf bytes.Buffer // Buffer initialization
	buf.Write([]byte("Hello "))
	fmt.Fprintf(&buf, "world!")
	buf.WriteTo(os.Stdout)

	fmt.Printf("\n\nDefer \n\n")
	fmt.Printf("Defer make function work first in last out\n\n")

	defer fmt.Printf("Hello-1\n") // will print last
	fmt.Printf("World-2\n")       // will print first
	defer defer_function()        // will print second
}
