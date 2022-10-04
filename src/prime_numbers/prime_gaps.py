# will calculate P_n+1 - P_n
import sys

LIMIT = 10000
# higher the LIMIT, better the accuracy

def main():
    primes = [] # it'll store primes
    for i in range(1, LIMIT): # cheking for primes 
        if isPrime(i):
            primes.append(i) # adding to our array

    difference = [] # it'll store the differences
    for i in range(len(primes) - 1):
        diff = primes[i+1] - primes[i] # P_n+1 - P_n
        difference.append(diff)

    print("Max difference:", max(difference))

def isPrime(n):
    """
    checks for a factor of n between 2 and (n - 1),
    if a factor is found, return False,
    if not return True.
    """
    if n <= 0:
        # returning error
        sys.stderr.write("Better luck next time\n")
        sys.exit(1)
    if n == 1:
        return False # 1 is'nt prime. is it?
    if n == 2:
        return True
    for i in range(2, n-1):
        if n % i == 0: # checking for factors
            return False
    return True # if prime return true

if __name__ == "__main__":
    main()
