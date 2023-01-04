# golang Composite Types

# Non-Reference Types

# (1) Arrays
### They are fixed-length sequences of the same type
### When you assign an array to another variable, it copies the entire array
### When you pass an array as an argument to a function, it makes an entire copy of the array
### using array will be better if you know the capacity of the array


## example
```go
	var arr1 [3]bool // size is 3 bytes

	arr1[0] = false
	arr1[1] = true
	arr1[2] = false

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%t\n", arr1[i]) // return false, true, false
	}
```


# (2) Structs

### structs is collection of fields. These fields can be of different type (object)

```go
	type dog struct {
		name string
		age  int8
	}

	var dog1 dog = dog{ // size is 24 bytes
		name: "John",
		age:  21,
	}

	fmt.Printf("dog name %s and age %d", dog1.name, dog1.age)  // dog name John and age 21
```

# ---End Non-Reference Types---


# Reference Types

# (1) Slices

### Slices are dynamically sized
### you can create slices by make([]TYPE, length, capacity)
### A slice does not store any data, it just describes a section of an underlying array.
### you can add item to slices by append built in function


### example
```go
	slices2 := []uint8{} // size is 24 bytes

	var arr2 [4]uint8
	arr2[0] = 41
	arr2[1] = 82
	arr2[2] = 12
	arr2[3] = 21

	slices2 = arr2[0:4]

	slices2 = append(slices2, 65)
	fmt.Println(slices2) // return [41 82 12 21 65]
```

## create slice by (make)
```go
	var strSlice = make([]string, 3)

	strSlice[0] = "Hello"
	strSlice[1] = " "
	strSlice[2] = "World"

	fmt.Println(reflect.ValueOf(strSlice).Kind())  // slice
    fmt.Printf("\n\nSlice Is %s\n\n", strings.Join(strSlice, "")) // Slice Is Hello World
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))  // strSlice Len: 3  Cap: 3
```

# (2) Map

### A map maps keys to values.

### example
```go
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

    fmt.Printf("person name is %s and age is %d\n\n", m["Foo"].name, m["Foo"].age) // person name is Foo and age is 24
	fmt.Printf("person name is %s and age is %d\n\n", m["Joe"].name, m["Joe"].age) // person name is Joe and age is 31

```
