package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(exp(1))
}

func exp(x float64) float64 {
    n := 2147483647.0 // highest value of float64
    sum := math.Pow((1 + x/n), n)
    return sum
}
