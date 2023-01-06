package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// Type assertions
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// Type switch
	do(21)
	do("hello")
	do(true)

	// interfaces
	GetAnimal()
}

type Animal interface {
	Feed(amount float32) float32
	Move(space uint32) uint32
}

type Dog struct {
	name       string
	age        uint32
	foodAmount float32
}

type Cat struct {
	name       string
	age        uint32
	foodAmount float32
}

func (dog *Dog) Feed(amount float32) float32 {
	return (dog.foodAmount * 4) * amount
}

func (dog *Dog) Move(space uint32) uint32 {
	return space - dog.age
}

func (cat *Cat) Feed(amount float32) float32 {
	return cat.foodAmount * amount
}

func (cat *Cat) Move(space uint32) uint32 {
	return (space * 4) - cat.age
}

func NewDog() Animal {
	return &Dog{
		name:       "Dogy",
		age:        4,
		foodAmount: 122,
	}
}

func NewCat() Animal {
	return &Cat{
		name:       "cate",
		age:        1,
		foodAmount: 40,
	}
}

func GetAnimal() {
	cat := NewCat()
	dog := NewDog()

	fmt.Printf("dog Feed: %v\n", dog.Feed(1221))
	fmt.Printf("dog Move: %v\n", dog.Move(232))
	fmt.Printf("cat Feed: %v\n", cat.Feed(112))
	fmt.Printf("cat Move: %v\n", cat.Move(8778))
}
