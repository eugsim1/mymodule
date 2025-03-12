// main.go
package main

import (
    "fmt"
    "github.com/eugsim1/mymodule"
)

func main() {
    message := mymodule.Greet("World")
    fmt.Println(message)
}
