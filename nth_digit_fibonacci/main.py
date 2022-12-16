# import math
#
#
# def get_is_primes(_max_size):
#     is_primes = [True for _ in range(_max_size + 1)]
#     primes = []
#     for i in range(2, _max_size):
#         if i * i >= _max_size:
#             break
#         if is_primes[i]:
#             for j in range(i * i, _max_size, i):
#                 is_primes[j] = False
#
#     for i in range(2, _max_size):
#         if is_primes[i]:
#             primes.append(i)
#
#     return is_primes, primes
#
#
# t = int(input().strip())
# inputs = []
# for a0 in range(t):
#     n = int(input().strip())
#     inputs.append(n)
#
# max_digit = max(inputs)
# max_size = pow(10, max_digit)
# is_primes, primes = get_is_primes(max_size)
#
# for ip in inputs:
#     min_for_ip = pow(10, ip-1)
#     max_for_ip = pow(10, ip)
#     for i in range(min_for_ip, max_for_ip):
#         if is_primes[i+1]:
#             print(primes.index(i+1))
#             break
import math


def get_fibonacci(_max_size):
    iter_count = 1
    while True:
        # ith_fibo = round(pow((1 + math.sqrt(5))/2, iter_count) / math.sqrt(5))
        # fibo_digit_count = int(math.log(ith_fibo, 10) + 1)
        # if iter_count < 10:
        # fibo = round(((1 + math.sqrt(5)) / 2) ** iter_count / math.sqrt(5))
        # fibo_digit_count = int(math.log10(fibo) + 1)
        # # else:
        fibo_digit_count = round(iter_count * math.log10((1 + math.sqrt(5)) / 2) - 0.5 * math.log10(5))
        iter_count += 1
        if fibo_digit_count == _max_size:
            return iter_count


t = int(input().strip())
inputs = []
for a0 in range(t):
    n = int(input().strip())
    print(get_fibonacci(n))
#
# max_digit = max(inputs)
# max_size = pow(10, max_digit)
# fibonaccies=
#
# print(fibonaccies)
# for ip in inputs:
#     min_for_ip = pow(10, ip - 1)
#     max_for_ip = pow(10, ip)
#     for i in range(min_for_ip, max_for_ip):
#         if is_primes[i + 1]:
#             print(primes.index(i + 1))
#             break
