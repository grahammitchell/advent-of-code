#!/usr/bin/python3

# 30 minutes

with open("input/10", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

def matches(c, m):
	return (c == '}' and m == '{') or (c == ']' and m == '[') or (c == ')' and m == '(') or (c == '>' and m == '<')

pointmap = {
	')': 1,
	']': 2,
	'}': 3,
	'>': 4
}

matching = {
	'{': '}',
	'[': ']',
	'(': ')',
	'<': '>'
}

scores = []
for line in puzzle:
	corrupt = False
	stack = []
	for c in line:
		if c in ( '{', '[', '<', '(' ):
			stack.append(c)
		elif c in ( '}', ']', '>', ')' ):
			m = stack.pop()
			if not matches(c, m):
				corrupt = True
				break
	if corrupt:
		continue
	# incomplete
	completion = ''
	points = 0
	while stack:
		m = stack.pop()
		c = matching[m]
		completion += c
		points = (points*5) + pointmap[c]
	scores.append(points)

scores.sort()
mid = len(scores)//2
print(scores[mid])
