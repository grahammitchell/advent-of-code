#!/usr/bin/python3

# 34 minutes

with open("input/13", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

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

max_x = 0
max_y = 0
xs = []
ys = []
for line in puzzle:
	if line.startswith("fold along"):
		first_fold = line
		break
	
	if ',' in line:
		(x,y) = [int(s) for s in line.split(",")]
		if x > max_x:
			max_x = x
		if y > max_y:
			max_y = y
		xs.append(x)
		ys.append(y)

(dir_s, val) = first_fold.split('=')
if dir_s.endswith('x'):
	xr = int(val)
	print(f"Fold on x={xr}")
	reflect(xs, xr)
if dir_s.endswith('y'):
	yr = int(val)
	print(f"Fold on y={yr}")
	reflect(ys, yr)

uniq = count_uniques(xs, ys)

print(uniq)
