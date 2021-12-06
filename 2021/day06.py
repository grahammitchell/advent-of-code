#!/usr/bin/python3

# 13 minutes

with open("input/6", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

fish = [int(s) for s in puzzle[0].split(",")]

day = 0
while day < 80:
	new_fish = 0
	for i,f in enumerate(fish):
		if f == 0:
			new_fish += 1
			fish[i] = 7
		fish[i] -= 1

	for i in range(new_fish):
		fish.append(8)
		
	day += 1
	print(fish)

print(len(fish))
