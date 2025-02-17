```go
package main

import "fmt"

func main() {
    var m map[string]int
    value, ok := m["a"]
    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("Key not found")
    }
}
```