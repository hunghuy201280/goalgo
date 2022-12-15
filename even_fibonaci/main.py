
t = int(input().strip())
for a0 in range(t):
	n = int(input().strip())
	# first two terms
	n1, n2 = 0, 1
	result = 0
	while n1 < n:
		nth = n1 + n2
		# update values
		n1 = n2
		n2 = nth
		if n1 >= n:
			break
		if n1 % 2 == 0:
			result += n1
	print(str(result) + "\n")