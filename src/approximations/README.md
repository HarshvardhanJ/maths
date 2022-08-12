## Approximations

This sub-section contains approximations to famous
constants.

### Pi

**Spigot algorithms**:

Two algorithms were discovered in 1995 that opened up new avenues of research into $π$. They are called [spigot algorithms](https://en.m.wikipedia.org/wiki/Spigot_algorithm) because, like water dripping from a [spigot](https://en.m.wikipedia.org/wiki/Tap_(valve)), they produce single digits of π that are not reused after they are calculated. This is in contrast to infinite series or iterative algorithms, which retain and use all intermediate digits until the final result is produced.


Another spigot algorithm, the [BBP](https://en.m.wikipedia.org/wiki/Bailey%E2%80%93Borwein%E2%80%93Plouffe_formula) [digit extraction algorithm](https://en.m.wikipedia.org/wiki/Digit_extraction_algorithm), was discovered in 1995 by Simon Plouffe:

$$ {\displaystyle \pi =\sum_{k=0}^{\infty}{\frac {1}{16^{k}}}\left({\frac {4}{8k+1}}-{\frac {2}{8k+4}}-{\frac {1}{8k+5}}-{\frac {1}{8k+6}}\right)} $$

**Implementation**

1. Python

```py
import math

def main():
    """main function"""
    times = 100.0
    # more times more accuracy, but too much is neglible
    sum = 0.0
    k = 0.0
    while k < times:
        mul = 1/16**k
        re = (4/(8*k + 1)) - (2/(8*k + 4)) - (1/(8*k + 5)) - (1/(8*k + 6))
        sum += mul*re
        k += 1
    print(sum)
    print(math.pi == sum) # => True, we would have exact value

if __name__ == "__main__":
    main()
```

2. Golang

```go
func main() {
    // driver code
    fmt.Println(bbp())
    fmt.Println(math.Pi)
    fmt.Println(bbp() == math.Pi) // => True, we would have exact value.
}

func bbp() float64 {
    // source: https://en.m.wikipedia.org/wiki/Pi#spigot-algorithms
    sum := 0.0
    for k := 0.0; k < 100.0; k++ {
        sum += (1/(math.Pow(16, k)))*(4/(8*k + 1) - 2/(8*k + 4) - 1/(8*k + 5) - 1/(8*k + 6))
    }
    return sum
}
```

### Speed of light

In [classical physics](https://en.m.wikipedia.org/wiki/Classical_physics), light is described as a type of [electromagnetic wave](https://en.m.wikipedia.org/wiki/Electromagnetic_wave). The classical behaviour of the [electromagnetic field](https://en.m.wikipedia.org/wiki/Electromagnetic_field) is described by [Maxwell's equations](https://en.m.wikipedia.org/wiki/Maxwell%27s_equations), which predict that the speed c with which electromagnetic waves (such as light) propagate in vacuum is related to the distributed capacitance and inductance of vacuum, otherwise respectively known as the [electric constant](https://en.m.wikipedia.org/wiki/Electric_constant) $\varepsilon _{0}$ and the [magnetic constant](https://en.m.wikipedia.org/wiki/Magnetic_constant) $\mu _{0}$, by the equation.

$$\displaystyle c={\frac {1}{\sqrt{\varepsilon _{0}\mu _{0}}}\.}$$

**Implementation**

1. Python

```py
import math

def main():
    """main function"""
    u_0 = 4*math.pi*math.pow(10, -7)
    e_0 = 8.8541878128 * math.pow(10, -12)
    c = 1/math.sqrt(u_0 * e_0)
    print(c) # => 299792458.0816064 

if __name__ == "__main__":
    main()
```

2. Golang

```go
func main() {
    u_0 := 4*math.Pi*math.Pow(10, -7)
    e_0 := 8.8541878128 * math.Pow(10, -12)
    c := 1/(math.Sqrt(u_0*e_0))
    fmt.Println(c) // => 2.997924580816064e+08
}
```
