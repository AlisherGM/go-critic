Detects when predeclared identifiers shadowed in assignments.

**Before:**
```go
func main() {
    // shadowing len function
    len := 10
    println(len)
}
```

**After:**
```go
func main() {
    // change identificator name
    length := 10
    println(length)
}
```