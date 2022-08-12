import math

def main():
    """main function"""
    u_0 = 4*math.pi*math.pow(10, -7)
    e_0 = 8.8541878128 * math.pow(10, -12)
    c = 1/math.sqrt(u_0 * e_0)
    print(c)

if __name__ == "__main__":
    main()
