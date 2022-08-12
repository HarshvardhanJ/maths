package main

import (
    "math"
    "fmt"
    "example.com/helper"
)

func main() {
    // driver code
    x := sin(math.Pi/3)
    y := cos(math.Pi/3)  
    fmt.Println(helper.Pow(x, 2) + helper.Pow(y, 2))
}

func sin(x float64) float64 {
    sum := 0.0
    for n := 0.0; n < 1000.0; n++ {
        re := (math.Pow(-1, n))/(Factorial(2*n + 1))
        y := math.Pow(x, 2*n + 1)
        sum += re*y
    }
    return sum
}

func cos(x float64) float64 {
    sum := 0.0
    for n := 0.0; n < 1000.0; n++ {
        re := (math.Pow(-1, n))/(Factorial(2*n))
        y := math.Pow(x, 2*n)
        sum += re*y
    }
    return sum
}

func Factorial(n float64) float64 {
    if n == 0 { return 1 }
    return Factorial(n-1)*n
}
