#!/usr/bin/python3

# 9 minutes

with open("input/8", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

outputs = []
for line in puzzle:
	half = line.split("|")[1].strip()
	outputs.extend( half.split() )

count = 0
for o in outputs:
	if len(o) in (2,3,4,7):
		count += 1

print(count)
