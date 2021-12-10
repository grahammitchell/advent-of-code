#!/usr/bin/python3

# 15 minutes

with open("input/10", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

def matches(c, m):
	return (c == '}' and m == '{') or (c == ']' and m == '[') or (c == ')' and m == '(') or (c == '>' and m == '<')

pointmap = {
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137
}

points = 0
for line in puzzle:
	stack = []
	for c in line:
		if c in ( '{', '[', '<', '(' ):
			stack.append(c)
		elif c in ( '}', ']', '>', ')' ):
			m = stack.pop()
			if not matches(c, m):
				points += pointmap[c]
				print(c)
				break

print(points)
