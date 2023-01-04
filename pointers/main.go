package main

import "fmt"

func i_take_pointer(num *int8) {
	*num = 80 - *num
}

func main() {

	fmt.Printf("Use of Pointers is to get the most performers out of a program\n\n")
	fmt.Printf("when u pass argument to a function it make another copy of it\n\n")
	fmt.Printf("but if u use pointer as argument it will use the same variable\n\n")
	fmt.Printf("use &variable to get memory address\n\n")
	fmt.Printf("use *pointer to extract value from this memory address\n\n")

	var num int8 = 122 // num memory address is 0xc000010148

	var ptr *int8 = &num // now ptr equaled to 0xc000010148

	fmt.Printf("num is %d and num memory address is %p\n\n", num, ptr) // return num is 122 and num memory address is 0xc000010148
	fmt.Printf("pointer value is %d\n\n", *ptr)                        // pointer value is 122

	fmt.Printf("num before %d\n\n", num) // num before 122
	i_take_pointer(&num)                 // passing 0xc000010148 to i_take_pointer function
	fmt.Printf("num after %d\n\n", num)  // num after -42
}
