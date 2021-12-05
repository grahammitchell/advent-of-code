#!/usr/bin/python3

# 26 minutes

with open("input/5", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

grid = []
for i in range(1000):
	grid.append( [0 for x in range(1000)] )

def gen_ys(x1, y1, x2, y2):
	ys = []
	for y in range(y1,y2+1):
		ys.append(y)
	return [(x1, y) for y in ys]

def gen_xs(x1, y1, x2, y2):
	xs = []
	for x in range(x1,x2+1):
		xs.append(x)
	return [(x, y1) for x in xs]

def gen_points(start, end):
	(x1, y1) = start
	(x2, y2) = end
	if x1 == x2:
		if y1 > y2:
			# put them in ascending order
			t = y1
			y1 = y2
			y2 = t
		return gen_ys(x1,y1,x2,y2)
	else:
		if x1 > x2:
			# put them in ascending order
			t = x1
			x1 = x2
			x2 = t
		return gen_xs(x1,y1,x2,y2)

def tally(points):
	for p in points:
		(x, y) = p
		grid[x][y] += 1

starts = []
ends = []
for line in puzzle:
	(a0, b0) = line.split(' -> ')
	(x1, y1) = a0.split(',')
	(x2, y2) = b0.split(',')
	if x1 != x2 and y1 != y2:
		continue
	starts.append((int(x1),int(y1)))
	ends.append((int(x2),int(y2)))

for i in range(len(starts)):
	start = starts[i]
	end = ends[i]
	points = gen_points(start, end)
	tally(points)

total = 0
for r in range(1000):
	for c in range(1000):
		if grid[r][c] >= 2:
			total += 1
print(total)
