package main

import (
    "fmt"
    "math"
)

func main() {
    sum := 0.0
    for i := 1.0; i < 10000000000; i++ {
        sum += 1/(i*i)
    }
    y := math.Pow(math.Pi, 2)
    fmt.Println("Final sum =>", sum, y/6)
}
