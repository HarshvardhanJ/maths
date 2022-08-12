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
