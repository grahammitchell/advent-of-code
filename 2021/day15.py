#!/usr/bin/python3

# ?? minutes

with open("input/15s", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

risk = []
for line in puzzle:
	risk.append( [int(s) for s in line] )

ROWS = len(risk)
COLS = len(risk[0])

r = ROWS-1
c = COLS-1

total = 0
cache = []
for line in puzzle:
	cache.append( [0 for s in line] )

visited = []
for line in puzzle:
	visited.append( [False for s in line] )

for r in range(ROWS):
	for c in range(COLS):
		total += risk[r][c]
maxx = total

def display(r,c,up,dn,lf,rt):
	if up < dn and up < lf and up < rt:
		print(f"({r-1},{c})")
	elif dn < up and dn < lf and dn < rt:
		print(f"({r+1},{c})")
	elif lf < up and lf < dn and lf < rt:
		print(f"({r},{c-1})")
	elif rt < up and rt < dn and rt < lf:
		print(f"({r},{c+1})")

def getBestPath(r,c, dest_r, dest_c,visited):
	if r == dest_r and c == dest_c:
		return risk[r][c]
	if cache[r][c]:
		return cache[r][c]
	if visited[r][c]:
		return maxx
	visited[r][c] = True
	up = dn = lf = rt = maxx
	if r+1 < ROWS:
		dn = getBestPath(r+1,c, dest_r, dest_c, visited)
	if c+1 < COLS:
		rt = getBestPath(r,c+1, dest_r, dest_c, visited)
	if r-1 >= 0:
		up = getBestPath(r-1,c, dest_r, dest_c, visited)
	if c-1 >= 0:
		lf = getBestPath(r,c-1, dest_r, dest_c, visited)
	best = risk[r][c] + min(up,dn,lf,rt)
	cache[r][c] = best
	return best

def getBestPath2(risk):
	sol = []
	for line in puzzle:
		sol.append([0 for s in line])

	sol[0][0] = risk[0][0]
	for i in range(1,COLS):
		sol[0][i] = risk[0][i] + sol[0][i-1]
	for i in range(1,ROWS):
		sol[i][0] = risk[i][0] + sol[i-1][0]

	for i in range(1,ROWS):
		for j in range(1, COLS):
			sol[i][j] = risk[i][j] + min(sol[i-1][j], sol[i][j-1])

	return sol[ROWS-1][COLS-1]
	
total = getBestPath(0,0,ROWS-1,COLS-1,visited)
#total = getBestPath2(risk)
total -= risk[0][0]
print(total)
