# golang Basic use

# (1) For Loop

### example
```go
	for i := 0; i < 10; i++ {
		fmt.Printf("\n i: %d", i)
	}
```

# (2) Range

### Range is a form of for loop that can be used with slices or maps

### example
```go
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2*%d = %d\n", i, v)
	}
```

# (3) If Statement

### example
```go
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

```


# (4) Switch Case

### example
```go
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
```


# (5) Error handling

### they are no try catch in golang

### example
```go
	src, err := os.ReadFile("./text.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("file content is \"%s\"", strings.TrimSpace(string(src)))
	}
```


# (6) Panic and Recover

### Panic will disable any function that bellow it when it called
### Panic will not disable deferred function that in top of it
### Panic will run deferred function after panic massage print
### Go provides us with the ability to regain control after panic has occurred

### example
```go
type user struct {
	name    string
	isAdmin bool
	id      int32
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
}

```

# (7) Defer

### Defer make function work (first in last out)

### example
```go
func defer_function() {
	defer fmt.Printf("World-3\n")
}

func main() {
    defer fmt.Printf("Hello-1\n") // will print last
	fmt.Printf("World-2\n")       // will print first
	defer defer_function()        // will print second
}
```
## ----------------End Of Basic Use-------------
