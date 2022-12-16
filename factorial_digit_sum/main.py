import math

t = int(input().strip())
for i in range(0, t):
    n = int(input().strip())

    factorial_n=math.factorial(n)
    num_str = str(factorial_n)

    res = 0

    for char in num_str:
        digit = int(char)
        res+=digit

    print(res)
