import math

def main():
    print(exp(1))

def exp(x):
    n = 2147483647 # max value of integer, we'll treat this as infinity.
    sum = math.pow((1 + x/n), n)
    return sum

if __name__ == "__main__":
    main()
