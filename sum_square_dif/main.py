#!/bin/python3

import sys


t = int(input().strip())
for a0 in range(t):
	n = int(input().strip())
	sum1=n*(n + 1)*(2*n + 1) / 6
	sum2=(n+1)*n/2

	print(int(abs(sum2*sum2-sum1)))
