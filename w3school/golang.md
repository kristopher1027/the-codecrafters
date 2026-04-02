# W3SCHOOL GOLANG 

## 1. Go Home

**Concept:** Entry point to Go language overview\
- Introduces Go as a compiled, statically typed language\
- Focus: simplicity, speed, concurrency

## Important Notes
* Used in backend systems, APIs, cloud tools
* Compiles to machine code → very fast



## 2. Go Introduction

**Concept:** What Go is and why it exists\
- Created at Google to solve slow builds + complex systems\
- Simple syntax, fast compilation, built-in concurrency
### Important Notes
* Competes with C, Java, Python (backend)


## 3. Go Get Started

``` go
package main
import "fmt"

func main() {
    fmt.Println("Hello World")
}
```
### Explanation
* package main → executable program
* main() → starting point
* fmt → standard output package

## 4. Go Syntax
### Concept: Basic structure rules
### key rules
-   No semicolons needed
-   {} required
-   Case-sensitive

## 5. Go Comments
### Concept: Add explanations in code
### Important Notes
* Ignored by compiler
* Used for readability

// single line\
/\* multi-line \*/



## 6. Go Variables
### Concept: Store data values

var x int = 10
y := 20

**Explanation**

* var → explicit
* := → shorthand (type inferred)

**Important Notes**

* := only works inside functions



## 7. Go Constants
### Concept: Fixed values (cannot change)

const pi = 3.14
### Important Notes
* Must be known at compile time
* Cannot reassign



## 8. Go Output
### Concept: Display output to user

fmt.Println("Hello")
fmt.Print("hi")
### Explanation
* Println → adds newline
* Print → no newline



## 9. Go Data Types
### Concept: Types of values
int, float64, string, bool
*example*
```go
var name string = "Go"
```
### Important Notes
* Go is strongly typed



## 10. Go Arrays
### Concept: Fixed-size collection
```go
var arr = [3]int{1,2,3}
```
### Explanation

* Size must be known
###  Important Notes
* Not flexible → rarely used directly



## 11. Go Slices
### Concept: Dynamic arrays
```go
s := []int{1,2,3}
```

### Explanation
* Can grow/shrink
### Important Notes
* Most used collection in Go


## 12. Go Operators
### Concept: Perform operations
**Types**

* Arithmetic: + - * /
* Comparison: == !=
* Logical: && ||

*Example*
```go
x := 5 + 3
```



## 13. Go Conditions

### Concept: Decision making

*Syntax*
```go
if x > 10 {
    fmt.Println("Big")
} else {
    fmt.Println("Small")
}
```
### Important Notes
* No parentheses needed


## 14. Go Switch
### Concept: Multiple conditions
*Syntax*
```go
switch x {
case 1:
    fmt.Println("One")
default:
    fmt.Println("Other")
}
```

### Important Notes
* Cleaner than many if-else



## 15. Go Loops
### Concept: Repetition
*Syntax*
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
### Explanation
* only loop in Go = for
### But it can act as:
* normal loop
* while loop
* infinite loop


> range
Used to loop over collections (array, slice, map)
```go
for i, v := range arr {
    fmt.Println(i, v)
}
```
> Nested loops
Yes, Go supports loops inside loops:
```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Println(i, j)
    }
}
```



## 16. Go Functions
### Concept: Reusable blocks of code
*Syntax*
```go
func add(a int, b int) int {
    return a + b
}
```
### Important Notes
* Can return multiple values



## 17. Go Struct
### Concept: Custom data type
*Syntax*
```go
type Person struct {
    name string
    age  int
}
```
### Explanation
* Groups related data



## 18. Go Maps
### Concept: Key-value storage
*Syntax*
```go
m := map[string]int{
    "a": 1,
}
```
### Explanation
* Fast lookup using keys
### Important Notes
* Similar to dictionary in other languages