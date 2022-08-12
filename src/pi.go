package main

import (
    "fmt"
    "math"
)

func main() {
    // driver code
    fmt.Println(bbp())
    fmt.Println(math.Pi)
    fmt.Println(bbp() == math.Pi)
}

func bbp() float64 {
    // source: https://en.m.wikipedia.org/wiki/Pi#spigot-algorithms
    sum := 0.0
    for k := 0.0; k < 100.0; k++ {
        sum += (1/(math.Pow(16, k)))*(4/(8*k + 1) - 2/(8*k + 4) - 1/(8*k + 5) - 1/(8*k + 6))
    }
    return sum
}

