#!/usr/bin/python3

# 20 minutes

with open("input/9", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

def is_low_point(grid, r, c):
	cur = grid[r][c]
	return cur < grid[r-1][c] and cur < grid[r][c-1] and cur < grid[r][c+1] and cur < grid[r+1][c]

	#return cur < grid[r-1][c-1]
	#	and cur < grid[r-1][c]
	#	and cur < grid[r-1][c+1]
	#	and cur < grid[r][c-1]
	#	and cur < grid[r][c+1]
	#	and cur < grid[r+1][c-1]
	#	and cur < grid[r+1][c]
	#	and cur < grid[r+1][c+1]

heights = []
for line in puzzle:
	heights.append( [int(s) for s in list(line)] )

# buffer with 10s
for row in heights:
	row.insert(0, 10)
	row.append(10)
heights.insert(0, [10]*len(heights[0]))
heights.append([10]*len(heights[0]))

risk = 0
for r in range(1,len(heights)):
	for c in range(1, len(heights[0])):
		if is_low_point(heights, r, c):
			risk += 1+heights[r][c]
print(risk)
