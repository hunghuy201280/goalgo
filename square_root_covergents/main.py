#!/bin/python3

import sys

t = int(input().strip())

prefN=3
prefD=2
for i in range(2,t+1):
    curN=prefN+2*prefD
    curD=prefN + prefD
    prefN=curN
    prefD=curD
    if len(str(curN))> len(curD):
        print(i)
