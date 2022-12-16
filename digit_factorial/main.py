import math

# t = int(input().strip())
# for a0 in range(t):
#     n = str(input().strip())

t = int(input().strip())
res = 0
for i in range(10, t):
    cur_num = i
    num_str = str(i)
    cur_sum = 0
    for char in num_str:
        digit = int(char)
        cur_sum += math.factorial(digit)

    if cur_sum % cur_num == 0:
        res += cur_num

print(res)
