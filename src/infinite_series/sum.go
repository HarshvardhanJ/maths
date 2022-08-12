package main

import (
    "fmt"
    "math"
)

func main() {
    times := 1000.0
    sum := 0.0
    for i := 1.0; i < times; i++ {
        sum += 1/(math.Pow(2, i))
        fmt.Println(sum, "=>", i)
        if sum == 1 {
            fmt.Println("1 at", i)
            break
        }
    }
}
