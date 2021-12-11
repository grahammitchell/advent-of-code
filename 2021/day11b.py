#!/usr/bin/python3

# 41 minutes

with open("input/11", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

grid = []
flashed = []
for line in puzzle:
	grid.append( [int(s) for s in line] )
	flashed.append( [False for s in line] )
ROWS = len(grid)
COLS = len(grid[0])

def reset():
	for r,row in enumerate(flashed):
		for c,octo in enumerate(flashed[r]):
			if flashed[r][c]:
				grid[r][c] = 0
			flashed[r][c] = False

def sum_of(grid):
	count = 0
	for r,row in enumerate(grid):
		for c,octo in enumerate(grid[r]):
			count += grid[r][c]
	return count

def count_flashes():
	flashes = 0
	for r,row in enumerate(flashed):
		for c,octo in enumerate(flashed[r]):
			if flashed[r][c]:
				flashes += 1
	return flashes

def flash(grid, r, c):
	if flashed[r][c]:
		return False
	
	# upper left
	if 0 <= r-1 and 0 <= c-1:
		grid[r-1][c-1] += 1
	# above
	if 0 <= r-1:
		grid[r-1][c] += 1
	# upper right
	if 0 <= r-1 and c+1 < COLS:
		grid[r-1][c+1] += 1
	
	# left
	if 0 <= c-1:
		grid[r][c-1] += 1
	# right
	if c+1 < COLS:
		grid[r][c+1] += 1

	# lower left
	if r+1 < ROWS and 0 <= c-1:
		grid[r+1][c-1] += 1
	# below
	if r+1 < ROWS:
		grid[r+1][c] += 1
	# lower right
	if r+1 < ROWS and c+1 < COLS:
		grid[r+1][c+1] += 1
	
	flashed[r][c] = True
	return True
	
def increment_errybody_by_one(grid):
	for r,row in enumerate(grid):
		for c,octo in enumerate(grid[r]):
			grid[r][c] += 1

def errybody_over_9_flashes(grid):
	for r,row in enumerate(grid):
		for c,octo in enumerate(grid[r]):
			if grid[r][c] > 9:
				flash(grid,r,c)

def do_step(grid):
	increment_errybody_by_one(grid)
	last = 1
	while True:
		cur = count_flashes()
		if cur == last:
			break
		errybody_over_9_flashes(grid)
		last = cur
	flashes = last
	reset()
	return flashes

step = 1
while True:
	flashes = do_step(grid)
	if flashes == ROWS*COLS:
		print(step)
		break
	step += 1

