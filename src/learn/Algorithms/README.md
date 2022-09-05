# Algorithms

This sub-section contains some algorithms related to maths that I am learning.

## Factors

Printing out all the factors of the given number.

**Implementation**

- Python

```py
./factor.py


#!/usr/bin/env python3
import sys

if len(sys.argv) != 2:
    # error msg
    sys.stderr.write("Usage: ./factor <number>\n")
    sys.exit(1)

# getting number from sys args
arg = sys.argv[1]

def main():
    """main function"""
    try:
        # converting string to int
        num_int = int(arg)
    except Exception as e:
        # error safe program
        print(e)
        sys.exit(1)
    # getting the list of factors in `re`
    re = factor(num_int)
    if len(re) == 0:
        # checking if factors are 0
        print(f"{num_int}: No factors found.")
    else:
        factors_str = ""
        # converting [factors] to "factors"
        for i in range(len(re)):
            if i == len(re) - 1:
                # not adding space in last element
                factors_str += str(re[i])
            else:
                factors_str += str(re[i]) + " "
        # printing out all the factors as a string
        print(f"{num_int}: {factors_str}")

def factor(n):
    # function to return an array of factors of n
    results = [] 
    for i in range(2, n):
        # ignoring 1 and n
        if n % i == 0:
            results.append(i)
    return results

if __name__ == "__main__":
    main()
```

- Golang

```go
./factor.go


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
```

---
