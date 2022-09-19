# Conjectures

In mathematics, a conjecture is a conclusion or a proposition that is proffered on a tentative basis without proof.

## Collatz Conjecture

The Collatz conjecture is one of the most famous unsolved problems in mathematics. The conjecture asks whether repeating two simple arithmetic operations will eventually transform every positive integer into 1. It concerns sequences of integers in which each term is obtained from the previous term as follows: if the previous term is even, the next term is one half of the previous term. If the previous term is odd, the next term is 3 times the previous term plus 1. The conjecture is that these sequences always reach 1, no matter which positive integer is chosen to start the sequence.

$${\displaystyle f(n)={\begin{cases}{\frac {n}{2}}&{\text{if }}n\equiv 0{\pmod {2}}3n+1&{\text{if }}n\equiv 1{\pmod {2}}.\end{cases}}}$$

**Implementation**

```c
// ./collatz_conjecture.c

#include <stdio.h>

// defining function
int checkEvenOrOdd(int num);

int main() {
    /* this conjecture states that 
     * every natural number will eventually lead to 1
     */

    /* Rules:
     * If Odd: n = 3n + 1
     * If Even: n = n/2
     */

    /* Height of tree:
     * Number of iterations throughout 1 is obtained
     */

    printf("Collatz Conjecture\n");
    printf("__________________\n\n");
    // sample number data
    int numbers[8] = {23, 34, 68, 23, 2048, 1, 0,10};
    for (int i = 0; i < 8; ++i) {
        int number = numbers[i];
        int height = 0;
        printf("Testing on %d\n", number);

        if (number < 1) {
            printf("Only numbers greater than 1 are supported.\n");
            return 1;
        }

        while (number != 1) {
            if (checkEvenOrOdd(number) == 1) {
                number = number / 2;
                ++height;
                printf("%d\n", number);
            } else {
                number = 3 * number + 1;
                ++height;
                printf("%d\n", number);
            }
        }
        printf("Height: %d\n\n", height);
    }
    return 0;
}

int checkEvenOrOdd(int num) {
    // checking if number is odd or even
    if (num % 2 == 0) {
        return 1;
    } else {
        return 0;
    }
}
```

---
