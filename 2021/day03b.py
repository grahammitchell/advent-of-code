#!/usr/bin/python3

with open("input/3", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

bits = len(puzzle[0])

def most_common(pos, nums):
	ones = sum([int(num[pos]) for num in nums])
	zeros = len(nums) - ones
	mcv = "1" if ones >= zeros else "0"
	lcv = "0" if mcv == "1" else "1"
	return (mcv, lcv)

oxy_list = list(puzzle)
for i in range(bits):
	(mcv, lcv) = most_common(i, oxy_list)
	oxy_list = [num for num in oxy_list if num[i] == mcv]
	if len(oxy_list) == 1:
		break
oxy = int(oxy_list[0], base=2)

co2_list = list(puzzle)
for i in range(bits):
	(mcv, lcv) = most_common(i, co2_list)
	co2_list = [num for num in co2_list if num[i] == lcv]
	if len(co2_list) == 1:
		break
co2 = int(co2_list[0], base=2)

print(f"{oxy}, {co2}")
result = oxy  * co2
print(result)
