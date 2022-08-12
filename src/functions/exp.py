import math

def main():
    print(exp(-1))
    print(exp(1))
    print(exp(30))

def exp(x):
    sum = 0.0
    times = 100.0
    i = 0.0
    while i < times:
        sum += (math.pow(x, i)) / factorial(i)
        i += 1
    return sum

def factorial(n):
    if n == 0: return 1
    return n*factorial(n-1)

if __name__ == "__main__":
    main()
