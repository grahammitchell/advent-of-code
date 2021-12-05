#!/usr/bin/python3

with open("input/3", "r") as f:
	puzzle = f.readlines()

diag = [int(s, base=2) for s in puzzle]

gamma = 0
epsilon = 0
for shift in range(12):
	count = [ 0, 0 ]
	for num in diag:
		digit = (num>>shift)&1
		count[digit] += 1
	(zeros, ones) = count
	cur = 1 if ones > zeros else 0
	other = 0 if ones > zeros else 1
	gamma += cur<<shift
	epsilon += other<<shift

print(f"{gamma}, {epsilon}")
result = gamma  * epsilon
print(result)
