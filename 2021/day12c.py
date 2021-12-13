#!/usr/bin/python3

# 60 minutes

import time

start = time.time()

with open("input/12", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

graph = {}
visited = {}
for line in puzzle:
	(a, b) = line.split('-')
	if a in graph:
		graph[a].append(b)
	else:
		graph[a] = [b]

	if b in graph:
		graph[b].append(a)
	else:
		graph[b] = [a]

	visited[a] = 0
	visited[b] = 0

def reset(visited):
	for k in visited:
		visited[k] = 0

def small_twice():
	for k in visited:
		if k.islower() and visited[k] > 1:
			return True
	return False

def can_visit(label):
	if label.isupper():
		return True
	elif label == 'start':
		return False
	elif visited[label] == 0 and label == 'end':
		return True
	elif small_twice() and visited[label] == 0:
		return True
	elif not small_twice() and visited[label] < 2:
		return True
	return False


def printAllPathsFrom(start, end, path, paths):
	visited[start] += 1
	path.append(start)
	if start == end:
		paths.append(",".join(path))
	else:
		for i in graph[start]:
			if can_visit(i):
				printAllPathsFrom(i, end, path, paths)
	path.pop()
	visited[start] -= 1


reset(visited)
path = []
paths = []
printAllPathsFrom('start', 'end', path, paths)
print(len(paths))

stop = time.time()
print(stop - start)
