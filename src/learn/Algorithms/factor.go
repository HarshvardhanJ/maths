package main

import (
    "fmt"
    "strconv"
    "os"
    "log"
)

func main() {
    log.SetFlags(0) // removing timestamp prefix
    if len(os.Args) != 2 {
        // error/help message
        log.Fatal("Usage: ./factor <number>")
    }
    text := os.Args[1] // getting text from sys args
    num, err := strconv.Atoi(text)
    if err != nil {
        log.Fatal(err)
    }
    re := factors(num) // calling factors
    if len(re) == 0 {
        fmt.Printf("%d: No factor found.\n", num)
    } else {
        factor_text := ""
        // converting [factors] to "factors"
        for n, i := range re {
            if n == len(re) - 1 {
                // last element
                // not adding extra whitespace here
                factor_text += strconv.Itoa(i)
            } else {
                factor_text += strconv.Itoa(i) + " "
            }
        }
        fmt.Printf("%d: %v\n", num, factor_text)
    }
}

func factors(n int) []int {
    // factors will return an array of all factors 
    var res []int
    for i := 2; i < n - 1; i++ {
        if n % i == 0 {
            // add factor to our array
            res = append(res, i)
        }
    }
    return res
}
