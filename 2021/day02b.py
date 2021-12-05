#!/usr/bin/python3

# 2021-12-02 

with open("input/2", "r") as f:
	commands = f.readlines()

horiz = 0
depth = 0
aim = 0

for line in commands:
	(cmd, xs) = line.split()
	x = int(xs)
	if cmd == "forward":
		horiz += x
		depth += aim*x
	elif cmd == "down":
		aim += x
	elif cmd == "up":
		aim -= x
	else:
		print(f"Unrecognized command: {cmd}")

print(f"{horiz}, {depth}")
result = horiz * depth
print(result)
