#!/bin/python3

import sys

t = int(input().strip())
MAX_N = 1000003
isPrimes = [True for _ in range(MAX_N + 1)]
isPrimes[0] = False
isPrimes[1] = False
i = 2
while True:
    if i * i >= MAX_N:
        break
    if isPrimes[i]:
        j = i * i
        while True:
            if j >= MAX_N:
                break
            isPrimes[j] = False
            j += i
    i += 1


prime_sums = [0 for _ in range(MAX_N + 1)]

i_prime = 0
cur_sum = 0
i_sum = 0
while True:
    if i_sum> MAX_N:
        break
    if isPrimes[i_sum]:
        cur_sum += i_sum
    prime_sums[i_sum]=cur_sum
    i_sum+=1

for a0 in range(t):
    n = int(input().strip())
    print(prime_sums[n])
