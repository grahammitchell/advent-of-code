#!/usr/bin/python3

# 42 minutes

with open("input/6", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

fish = [int(s) for s in puzzle[0].split(",")]

fish2 = [0,0,0,0,0,0,0,0,0]

for f in fish:
	fish2[f] += 1
print(fish2)

day = 0
while day < 256:
	new_fish = fish2[0]
	fish2[0] = fish2[1]
	fish2[1] = fish2[2]
	fish2[2] = fish2[3]
	fish2[3] = fish2[4]
	fish2[4] = fish2[5]
	fish2[5] = fish2[6]
	fish2[6] = fish2[7]
	fish2[7] = fish2[8]

	fish2[8] = new_fish
	fish2[6] += new_fish
	day += 1

print(sum(fish2))
