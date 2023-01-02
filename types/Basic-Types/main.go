package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// integers
	fmt.Printf("integers......\n\n")

	fmt.Printf("there signed numbers (int), and unsigned numbers (unit)\n\n")

	fmt.Printf("signed numbers (int) can hold positive or negative values \n\n")
	fmt.Printf("unsigned signed numbers (unit) can hold only positive values \n\n")

	fmt.Printf("in case of int8 can Hold from (-128 to 127)  \n\n")
	fmt.Printf("in case of unit8 can hold from (0 to 255)  \n\n")

	fmt.Printf("NOTE-- unit and int takes the same space in memory  \n\n")

	var num1 int8 = 127
	var num2 int16 = 12
	var num3 int32 = 2147483647
	var num4 int64 = 12
	var num5 uint8 = 255
	var num6 uint16 = 12
	var num7 uint32 = 12
	var num8 uint64 = 12

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num1))
	fmt.Printf("num's value is %d\n", num1)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num1))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num2))
	fmt.Printf("num's value is %d\n", num2)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num2))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num3))
	fmt.Printf("num's value is %d\n", num3)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num3))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num4))
	fmt.Printf("num's value is %d\n", num4)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num4))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num5))
	fmt.Printf("num's value is %d\n", num5)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num5))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num6))
	fmt.Printf("num's value is %d\n", num6)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num6))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num7))
	fmt.Printf("num's value is %d\n", num7)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num7))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(num8))
	fmt.Printf("num's value is %d\n", num8)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(num8))

	fmt.Printf("End Of Integers\n\n")

	// bool
	fmt.Printf("booleans......\n\n")
	var bool1 bool = true
	var bool2 bool = false

	fmt.Printf("bool's type is %s\n", reflect.TypeOf(bool1))
	fmt.Printf("bool's value is %t\n", bool1)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(bool1))

	fmt.Printf("bool's type is %s\n", reflect.TypeOf(bool2))
	fmt.Printf("bool's value is %t\n", bool2)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(bool2))

	fmt.Printf("End Of Booleans\n\n")

	// byte
	fmt.Printf("byte......\n\n")
	fmt.Printf("byte is the same as unit8 \n\n")
	fmt.Printf("byte is using \"American Standard Code for Information Interchange\"  \n\n")

	var byte1 byte = 97
	byte2 := []byte{97, 98, 99, 100, 101, 102}

	fmt.Printf("byte's type is %s\n", reflect.TypeOf(byte1))
	fmt.Printf("byte's value is %d\n", byte1)
	fmt.Printf("character equivalent to %d is \"%c\"\n", byte1, byte1)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(byte1))

	fmt.Printf("byte's type is %s\n", reflect.TypeOf(byte2))
	fmt.Printf("byte's value is %d\n", byte2)
	fmt.Printf("characters equivalent to %d is \"%c\"\n", byte2, byte2)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(byte2))

	fmt.Printf("End Of Byte\n\n")

	// float
	fmt.Printf("float......\n\n")
	fmt.Printf("the different between integer and float is that float can store decimals\n\n")

	var float1 float32 = 1242140000.503882901
	var float2 float64 = 52233223232233232232322222222.3232322332

	fmt.Printf("num's type is %s\n", reflect.TypeOf(float1))
	fmt.Printf("num's value is %f\n", float1)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(float1))

	fmt.Printf("num's type is %s\n", reflect.TypeOf(float2))
	fmt.Printf("num's value is %f\n", float2)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(float2))

	fmt.Printf("End Of Float\n\n")

	// rune
	fmt.Printf("rune......\n\n")
	fmt.Printf("rune is the same as int32 \n\n")
	fmt.Printf("rune is more powerful then byte but its use more memory \n\n")

	var rune1 rune = 1202

	rune2 := []rune{9221, 6121, 8877, 23422, 3343, 9778}

	fmt.Printf("rune's type is %s\n", reflect.TypeOf(rune1))
	fmt.Printf("rune's value is %d\n", rune1)
	fmt.Printf("character equivalent to %d is \"%c\"\n", rune1, rune1)
	fmt.Printf("size is %d runes\n\n", unsafe.Sizeof(rune1))

	fmt.Printf("rune's type is %s\n", reflect.TypeOf(rune2))
	fmt.Printf("rune's value is %d\n", rune2)
	fmt.Printf("characters equivalent to %d is \"%c\"\n", rune2, rune2)
	fmt.Printf("size is %d runes\n\n", unsafe.Sizeof(rune2))

	fmt.Printf("End Of Rune\n\n")

	// complex
	fmt.Printf("complex......\n\n")
	fmt.Printf("there are two parts of a complex The real and imaginary part made by float32 for complex64 and float64 for complex128\n\n")

	var complex1 complex64 = 10.32332 + 11i
	var complex2 complex128 = complex(10.32332, 4132)

	// getting real numbers
	r1 := real(complex1)
	r2 := real(complex2)

	// getting imaginary numbers
	i1 := imag(complex1)
	i2 := imag(complex2)

	fmt.Printf("complex's type is %s\n", reflect.TypeOf(complex1))
	fmt.Printf("complex's value is %f\n", complex1)
	fmt.Printf("complex's real part is %f\n", r1)
	fmt.Printf("complex's imaginary part is %f\n", i1)
	fmt.Printf("operations- \n addition: %f \n subtraction: %f \n multiplication: %f \n division: %f \n", r1+i1, r1-i1, r1*i1, r1/i1)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(complex1))

	fmt.Printf("complex's type is %s\n", reflect.TypeOf(complex2))
	fmt.Printf("complex's value is %f\n", complex2)
	fmt.Printf("complex's real part is %f\n", r2)
	fmt.Printf("complex's imaginary part is %f\n", i2)
	fmt.Printf("operations- \n addition: %f \n subtraction: %f \n multiplication: %f \n division: %f \n", r2+i2, r2-i2, r2*i2, r2/i2)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(complex2))

	fmt.Printf("End Of Complex\n\n")

	// string
	fmt.Printf("string......\n\n")

	var s string = "Hello World"
	fmt.Printf("string's type is %s\n", reflect.TypeOf(s))
	fmt.Printf("string's value is \"%s\"\n", s)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(s))

	fmt.Printf("End Of String\n\n")

	// uintptr
	fmt.Printf("uintptr......\n\n")
	fmt.Printf("uintptr is an integer representation of a memory address\n\n")

	var uintptr1 uintptr = 0xc80007b900

	fmt.Printf("uintptr's type is %s\n", reflect.TypeOf(uintptr1))
	fmt.Printf("uintptr's value is \"%d\"\n", uintptr1)
	fmt.Printf("size is %d bytes\n\n", unsafe.Sizeof(uintptr1))

	fmt.Printf("----------------End Of Basic Types-------------\n\n")
}
