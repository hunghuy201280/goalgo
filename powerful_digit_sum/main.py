#!/bin/python3

import sys

t = int(input().strip())
max = -1
for i in range(t):
    for j in range(t):
        temp = i ** j
        sum = 0
        while temp != 0:
            sum += temp % 10
            temp = temp // 10
        if sum > max:
            max = sum


print(max)
