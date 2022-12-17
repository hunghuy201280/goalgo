#!/bin/python3

import sys

t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    res = pow(2, n)
    sum = 0
    while res !=0:
        sum += res % 10
        res=res//10
    print(sum)
