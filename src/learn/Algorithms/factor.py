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
