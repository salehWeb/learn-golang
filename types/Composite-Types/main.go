package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	// Non-Reference Types
	fmt.Printf("Non-Reference Types......\n\n\n")

	fmt.Printf("Arrays\n\n")
	fmt.Printf("They are fixed-length sequences of the same type \n\n")
	fmt.Printf("When you assign an array to another variable, it copies the entire array\n\n")
	fmt.Printf("When you pass an array as an argument to a function, it makes an entire copy of the array\n\n")

	var arr1 [3]bool

	arr1[0] = false
	arr1[1] = true
	arr1[2] = false

	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(arr1))

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%t\n", arr1[i])
	}

	fmt.Printf("Structs\n\n")
	fmt.Printf("structs is collection of fields. These fields can be of different type\n\n")

	type dog struct {
		name string
		age  int8
	}

	var dog1 dog = dog{
		name: "John",
		age:  21,
	}

	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(dog1))

	fmt.Printf("dog name %s and age %d\n\n", dog1.name, dog1.age)

	fmt.Printf("End Non-Reference Types\n\n\n")

	// Reference Types
	fmt.Printf("Reference Types........\n\n\n")

	fmt.Printf("Slices\n\n")

	fmt.Printf("Slices are dynamically sized\n\n")
	fmt.Printf("you can create slices by make([]TYPE, length, capacity)\n\n")
	fmt.Printf("A slice does not store any data, it just describes a section of an underlying array.\n\n")
	fmt.Printf("you can add item to slices by append built in function\n\n")

	slices2 := []uint8{}

	var arr2 [4]uint8
	arr2[0] = 41
	arr2[1] = 82
	arr2[2] = 12
	arr2[3] = 21

	slices2 = arr2[0:4]

	slices2 = append(slices2, 65)
	fmt.Println(slices2)

	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(slices2))

	fmt.Printf("create slice by (make)\n\n")

	var strSlice = make([]string, 3)

	strSlice[0] = "Hello"
	strSlice[1] = " "
	strSlice[2] = "World"

	fmt.Println(reflect.ValueOf(strSlice).Kind())
	fmt.Printf("\n\nSlice Is %s\n\n", strings.Join(strSlice, ""))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

	fmt.Printf("\n\nMap\n\n")
	fmt.Printf("A map maps keys to values.\n\n")

	type Person struct {
		name string
		age  uint8
	}

	var m map[string]Person = map[string]Person{
		"Foo": {
			name: "Foo",
			age:  24,
		},
		"Joe": {
			name: "Joe",
			age:  31,
		},
	}

	fmt.Printf("person name is %s and age is %d\n\n", m["Foo"].name, m["Foo"].age)
	fmt.Printf("person name is %s and age is %d\n\n", m["Joe"].name, m["Joe"].age)

}
