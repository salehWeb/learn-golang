# golang Basic Types

# (1) integers

### there signed numbers (int), and unsigned numbers (unit)
### signed numbers (int) can hold positive or negative values
### unsigned signed numbers (unit) can hold only positive values

### in case of int8 can Hold from (-128 to 127)
### in case of unit8 can hold from (0 to 255)

```go
	var num1 int8 = 127 // size is 1 bytes
	var num2 int16 = 12 // size is 2 bytes
	var num3 int32 = 2147483647 // size is 4 bytes
	var num4 int64 = 12 // size is 8 bytes

	var num5 uint8 = 255 // size is 1 bytes
	var num6 uint16 = 12 // size is 2 bytes
	var num7 uint32 = 12 // size is 4 bytes
	var num8 uint64 = 12 // size is 8 bytes
```


### NOTE-- unit and int takes the same space in memory


# (2) booleans

### the default value of undeclared bool value is false

```go
	var bool1 bool = true // size is 1 bytes
	var bool2 bool = false // size is 1 bytes
    var bool3 bool
    fmt.Printf("%s", reflect.TypeOf(bool3)) // return false
```



# (3) byte

### byte is the same as unit8
### byte is using "American Standard Code for Information Interchange"


### example
```go
	var byte1 byte = 97 // size is 1 bytes as unit8
	byte2 := []byte{97, 98, 99, 100, 101, 102}

    fmt.Printf("\"%c\"", byte1) // return "a"
    fmt.Printf("\"%c\"", byte2) //  return "[a b c d e f]"

```


# (4) float

### the different between integer and float is that float can store decimals

### example
```go
	var float1 float32 = 1242140000.503882901 // size is 4 bytes
	var float2 float64 = 52233223232233232232322222222.3232322332 // size is 8 bytes
```


# (5) rune

### rune is the same as int32
### rune is more powerful then byte but its use more memory

### example
```go
	var rune1 rune = 1202 // size is 4 bytes as int32
	rune2 := []rune{9221, 6121, 8877, 23422, 3343, 9778}

    fmt.Printf("\"%c\"", rune1) // return  "Ҳ"
	fmt.Printf("\"%c\"", rune2) // return "[␅ ៩ ⊭ 孾 ഏ ☲]"
```


# (6) complex number

### there are two parts of a complex The real and imaginary part
### complex64 use two float32 and complex128 use two float64

### example
```go
	var complex1 complex64 = 10.32332 + 11i // shorted way. size is 8 bytes
	var complex2 complex128 = complex(10.32332, 4132) // size is 16 bytes

    // getting real numbers
	r1 := real(complex1) // 10.323320
	r2 := real(complex2) // 10.323320

	// getting imaginary numbers
	i1 := imag(complex1) // 11.000000
	i2 := imag(complex2) // 4132.000000

```


# (7) string

### You Know string 😂


### example
```go
	var s string = "Hello World" // size is 16 bytes
    fmt.Printf("\"%s\"", s) // return "Hello World"
```

### NOTE- strings are read only in golang

# (8) uintptr

### uintptr is an integer representation of a memory address

### example
```go
	var uintptr1 uintptr = 0xc80007b900 // size is 8 bytes
    fmt.Printf("\"%d\"", uintptr1)  // return "858993965312"
```
## ----------------End Of Basic Types-------------
