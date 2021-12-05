#!/usr/bin/python3

# 2021-12-01 

with open("input/1", "r") as f:
	words = f.readlines()
	measurements = [int(i) for i in words]

i = 1
prev = measurements[0]
increases = 0
while i < len(measurements):
	cur = measurements[i]
	if cur > prev:
		increases += 1
	prev = cur
	i += 1

print(increases)
