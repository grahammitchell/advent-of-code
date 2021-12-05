#!/usr/bin/python3

import copy, sys
from typing import NamedTuple

BLANK = [
	[ 0, 0, 0, 0, 0 ],
	[ 0, 0, 0, 0, 0 ],
	[ 0, 0, 0, 0, 0 ],
	[ 0, 0, 0, 0, 0 ],
	[ 0, 0, 0, 0, 0 ]
]

MARKED = [
	[ False, False, False, False, False ],
	[ False, False, False, False, False ],
	[ False, False, False, False, False ],
	[ False, False, False, False, False ],
	[ False, False, False, False, False ]
]

class Bingo:
	def __init__(self, board):
		self.board = copy.deepcopy(BLANK)
		self.marked = copy.deepcopy(MARKED)
		for i in range(5):
			line = board[i]
			self.board[i] = [int(s) for s in line.split()]
	
	def mark(self, n):
		for row in range(5):
			for col in range(5):
				if self.board[row][col] == n:
					self.marked[row][col] = True
					return True
		return False

	def is_winner(self):
		# check rows
		for row in range(5):
			if self.marked[row][0] and self.marked[row][1] and self.marked[row][2] and self.marked[row][3] and self.marked[row][4]:
				return True
		# check cols
		for col in range(5):
			if self.marked[0][col] and self.marked[1][col] and self.marked[2][col] and self.marked[3][col] and self.marked[4][col]:
				return True
		return False
	
	def unmarked(self):
		total = 0
		for row in range(5):
			for col in range(5):
				if not self.marked[row][col]:
					total += self.board[row][col]
		return total

boards = []

def everybody_won(already_won):
	for b in already_won.values():
		if not b:
			return False
	return True

with open("input/4", "r") as f:
	numberline = f.readline().strip()
	numbers = [int(s) for s in numberline.split(',')]
	boards_raw = f.readlines()

# initialize boards
while True:
	if not boards_raw:
		break
	board = boards_raw[1:6]
	b = Bingo(board)
	boards.append(b)
	boards_raw = boards_raw[6:]

# play bingo
last = 0
last_picked = 0

already_won = {}
for i in range(len(boards)):
	already_won[i] = False

for picked in numbers:
	print(picked)
	which = 0
	for b in boards:
		if everybody_won(already_won):
			continue
		b.mark(picked)
		if b.is_winner() and not already_won[which]:
			last = which
			last_picked = picked
			print(f"Setting last to {which}")
			already_won[which] = True
		which += 1

b = boards[last]
unmarked = b.unmarked()
print("BINGO!")
print(f"Board {last}")
print(f"{unmarked} {last_picked}")
result = unmarked * last_picked
print(result)

