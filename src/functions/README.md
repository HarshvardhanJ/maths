# Functions

This sub-section expansions/solutions to functions.

## Exponential function

**Graph:**

![Graph of e^x](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c6/Exp.svg/1600px-Exp.svg.png)

**Denoted by**: $f(x)=\exp(x)$ or $e^x$

**Domain**: $\mathbb {C}$

**Derivative**: ${\displaystyle \exp 'x=\exp x}$

**Series definition**: ${\displaystyle \exp x=\sum_{n=0}^{\infty }{\frac {x^{n}}{n!}}}$

**Implementation**

- Python

```py
# ./exp.py
def exp(x):
    sum = 0.0
    times = 100.0
    i = 0.0
    while i < times:
        sum += (math.pow(x, i)) / factorial(i)
        i += 1
    return sum
```

- Golang

```go
// ./exp.go
func exp(x float64) float64 {
    var sum float64
    var times float64 = 100.0
    for i := 0.0; i < times; i++ {
        temp := (Pow(x, i)) / Factorial(i)
        sum += temp
    }
    return sum
}
```

---

$${\displaystyle e^{x+y}=e^{x}e^{y}{\text{for all}x,y \in \mathbb {R},}}$$
which, along with the definition {\displaystyle e=\exp(1)}, shows that ${\displaystyle e^{n}=\underbrace{e\times \cdots \times e}_{n{\text{ factors}}}}$ for positive integers $n$, and relates the exponential function to the elementary notion of exponentiation.


[Euler's number](https://en.m.wikipedia.org/wiki/Euler%27s_number) $e$ = 2.71828... is the unique base for which the constant of proportionality is 1, since {\displaystyle \ln(e)=1}\ln(e)=1, so that the function is its own derivative:

$${\displaystyle {\frac {d}{dx}}e^{x}=e^{x}\ln(e)=e^{x}.}$$

<p align="center" width="100%">
    <img width="33%" src="https://upload.wikimedia.org/wikipedia/commons/thumb/6/62/Exp_series.gif/220px-Exp_series.gif" alt = "Exp series GIF">
</p>

> The exponential function (in blue), and the sum of the first $n + 1$ terms of its power series (in red).

The real exponential function {\displaystyle \exp \colon \mathbb {R} \to \mathbb {R} }{\displaystyle \exp \colon \mathbb {R} \to \mathbb {R} } can be characterized in a variety of equivalent ways. It is commonly defined by the following [power series](https://en.m.wikipedia.org/wiki/Power_series):

$${\displaystyle \exp x:=\sum _{k=0}^{\infty }{\frac{x^{k}}{k!}}=1+x+{\frac {x^{2}}{2}}+{\frac{x^{3}}{6}}+{\frac {x^{4}}{24}}+\cdots }$$

By way of the [binomial theorem](https://en.m.wikipedia.org/wiki/Binomial_theorem) and the power series definition, the exponential function can also be defined as the following limit:

$${\displaystyle \exp x=\lim _{n\to \infty }\left(1+{\frac {x}{n}}\right)^{n}.}$$

**Implementation**

- Python

```py
# ./exp1.py
def exp(x):
    n = 2147483647 # max value of integer, we'll treat this as infinity.
    sum = math.pow((1 + x/n), n)
    return sum
```

- Golang

```go
// ./exp1.go
func exp(x float64) float64 {
    n := 2147483647.0 // highest value of float64
    sum := math.Pow((1 + x/n), n)
    return sum
```

---
