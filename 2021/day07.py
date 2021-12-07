#!/usr/bin/python3

# 5 minutes

with open("input/7", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

crabs = [int(s) for s in puzzle[0].split(",")]
lo = min(crabs)
hi = max(crabs)

def calc_fuel(pos):
	fuel = 0
	for c in crabs:
		fuel += abs(c-pos)
	return fuel

best = hi*hi # easy to beat but in-bounds
for pos in range(lo, hi+1):
	cur = calc_fuel(pos)
	if cur < best:
		best = cur

print(best)
