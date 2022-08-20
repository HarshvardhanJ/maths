import math

i = 1
sum = 0
while i < 100:
    sum += (math.pow(-1, i+1)) / i
    i += 1
print(sum)
