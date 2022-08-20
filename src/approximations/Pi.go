package main

import (
    "fmt"
    "math"
)

func main() {
    sum := 0.0
    for i := 1.0; i < 1000.0; i++ {
        num := (math.Pow(-1.0, i+1.0)) * 4.0
        denom := 2.0*i - 1.0
        sum += num/denom
    }
    fmt.Println(sum)
}
