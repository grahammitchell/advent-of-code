#!/usr/bin/python3

# 43 minutes (total)

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

def get_mod(a, b):
	return int((b-a)/abs(b-a))

def gen_diag(x1, y1, x2, y2):
	result = [ (x1, y1) ]
	# 8,0 -> 0,8 - 7,1  6,2  5,3 
	# 6,4 -> 2,0: 5,3  4,2
	x_mod = get_mod(x1,x2)
	y_mod = get_mod(y1,y2)
	cur_x = x1
	cur_y = y1
	while True:
		x = cur_x + x_mod
		y = cur_y + y_mod
		result.append( (x,y) )
		cur_x = x
		cur_y = y
		if x == x2 and y == y2:
			break
	return result

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
	elif y1 == y2:
		if x1 > x2:
			# put them in ascending order
			t = x1
			x1 = x2
			x2 = t
		return gen_xs(x1,y1,x2,y2)
	else:
		test = gen_diag(x1,y1,x2,y2)
		return test

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
