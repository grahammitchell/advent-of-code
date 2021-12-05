#!/usr/bin/python3

# 2021-12-01 

with open("input/1", "r") as f:
	words = f.readlines()
	measurements = [int(i) for i in words]

i = 1
prev = sum(measurements[0:3])
increases = 0
while i < len(measurements):
	cur_slice = measurements[i:i+3]
	if len(cur_slice) < 3:
		break
	cur = sum(cur_slice)
	if cur > prev:
		increases += 1
	prev = cur
	i += 1

print(increases)
