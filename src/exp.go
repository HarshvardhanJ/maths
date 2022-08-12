package main

import (
    "fmt"
)

func main() {
    // driver code
    fmt.Println(exp(-1.0))
    fmt.Println(Pow(2, 6))
}

func exp(x float64) float64 {
    // TODO: Implement exp()
    var sum float64
    var times float64 = 100.0
    for i := 0.0; i < times; i++{
        temp := (Pow(x, i)) / Factorial(i)
        sum += temp
    }
    return sum
}

func Factorial(n float64) float64 {
    if n == 0 { return 1 }
    return n*Factorial(n-1)
}

func Pow(base, power float64) float64 {
    if base == 1 || power == 0 { return 1 }
    re := 1.0
    for i := 1.0; i <= power; i++{
        re *= base
    }
    return re
}
