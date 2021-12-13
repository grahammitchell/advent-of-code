#!/usr/bin/python3

# 46 minutes

with open("input/13", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

# read in input
xs = []
ys = []
for i,line in enumerate(puzzle):
	if line.startswith("fold along"):
		first_fold = line
		fold_loc = i
		break
	if ',' in line:
		(x,y) = [int(s) for s in line.split(",")]
		xs.append(x)
		ys.append(y)

def count_uniques(xs, ys):
	points = []
	for i,x in enumerate(xs):
		y = ys[i]
		points.append( (x,y) )
	point_set = set(points)
	return len(point_set)

def reflect(ys, yr):
	for i,y in enumerate(ys):
		if y > yr:
			y = yr*2 - y
			ys[i] = y

def get_maximums(xs, ys):
	max_x = 0
	max_y = 0
	for i,x in enumerate(xs):
		y = ys[i]
		if x > max_x:
			max_x = x
		if y > max_y:
			max_y = y
	return (max_x, max_y)

def visualize(xs, ys):
	(COLS, ROWS) = get_maximums(xs, ys)
	grid = [ ]
	for i in range(ROWS+1):
		grid.append( ['.' for j in range(COLS+1)] )
	for i,x in enumerate(xs):
		y = ys[i]
		grid[y][x] = '#'
	for r in range(ROWS+1):
		for c in range(COLS+1):
			print(grid[r][c], end='')
		print()


for fold in puzzle[i:]:
	(dir_s, val) = fold.split('=')
	if dir_s.endswith('x'):
		xr = int(val)
		print(f"Folding on x={xr}")
		reflect(xs, xr)
	if dir_s.endswith('y'):
		yr = int(val)
		print(f"Folding on y={yr}")
		reflect(ys, yr)

uniq = count_uniques(xs, ys)
visualize(xs, ys)

