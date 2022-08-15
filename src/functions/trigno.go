package main

import (
    "math"
    "fmt"
)

func main() {
    // driver code
    x := math.Pi/3
    fmt.Println(sin(x))
    fmt.Println(cos(x))
    fmt.Println(tan(x))
    fmt.Println(cot(x))
    fmt.Println(cosec(x))
    fmt.Println(sec(x))
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

func tan(x float64) float64 {
    return sin(x)/cos(x)
}

func cot(x float64) float64 {
    return cos(x)/sin(x)
}

func cosec(x float64) float64 {
    return 1/sin(x)
}

func sec(x float64) float64 {
    return 1/cos(x)
}

func Factorial(n float64) float64 {
    if n == 0 { return 1 }
    return Factorial(n-1)*n
}
