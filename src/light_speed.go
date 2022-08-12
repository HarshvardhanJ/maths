package main

import (
    "fmt"
    "math"
)

func main() {
    u_0 := 4*math.Pi*math.Pow(10, -7)
    e_0 := 8.8541878128 * math.Pow(10, -12)
    re := 1/(math.Sqrt(u_0*e_0))
    fmt.Println(re)
}
