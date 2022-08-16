import math

def main():
    x = 20
    data = send_fib(x)
    hash = {} # hashmap
    i = 0
    while i < len(data):
        hash[data[i]] = isPrime(data[i])
        i += 1
    primes = []
    for x in hash:
        print(x, "=>", hash[x])
        if hash[x] == True:
            primes.append(x)
    print(primes)

def send_fib(n):
    def fib(x):
        if x <= 1: return x
        return fib(x-1) + fib(x-2)
    array = []
    i = 0 
    while i < n:
        array.append(fib(i))
        i += 1
    return array

def isPrime(x):
    if x == 1 or x == 2: return True
    i = 2
    while i < x-1:
        if x % i == 0:
            return False
        i += 1
    return True

if __name__ == "__main__":
    main()
