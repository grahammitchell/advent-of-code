#!/usr/bin/python3

# 15 minutes

with open("input/7", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

crabs = [int(s) for s in puzzle[0].split(",")]
lo = min(crabs)
hi = max(crabs)

costs = [ 0 ]*(hi+1)
acc = 0
for i in range(hi+1):
	# I'm certain there's an equation for this but we're in a hurry
	acc += i
	costs[i] = acc

def calc_fuel(pos):
	fuel = 0
	for c in crabs:
		dist = abs(c-pos)
		fuel += costs[dist]
	return fuel

best = max(costs)*max(costs) # easy to beat but in-bounds
for pos in range(lo, hi+1):
	cur = calc_fuel(pos)
	if cur < best:
		best = cur

print(best)
