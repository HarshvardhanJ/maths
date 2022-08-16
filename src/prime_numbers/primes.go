package main

import (
    "fmt"
    "math"
)

func main() {
    // driver code
    var count int
    for i := 2; i < 1000; i++ {
        if isPrime(i){
            count += 1
        }
    }
    fmt.Println("0000 - 1000", "=>", count)
    for i := 1000; i < 10000; i = i + 1000 {
        count = 0
        for j := i; j < i + 1000; j++ {
            if isPrime(j){
                count += 1
            }
        }
        fmt.Println(i, "-", i+1000, "=>", count)
    }
    array := []int{10, 100, 1000, 10000, 100000}
    for _, i := range array {
        fmt.Println(prime(i))
    }
}

func isPrime(n int) bool {
    // TODO: Implement isPrime()
    for i := 2; i < n - 1; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func prime(n int) (int, float64) {
    // returns the amount of primes between 2 and n
    // TODO: Implement prime()
    list := []int{}
    for i := 2; i <= n; i++{
        if isPrime(i){
            list = append(list, i)
        }
    }
    m := float64(n)
    approx := (m/(math.Log(m)-1))
    return len(list), math.Floor(approx)
}
