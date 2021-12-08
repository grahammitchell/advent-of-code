#!/usr/bin/python3

# 72 minutes

with open("input/8", "r") as f:
	puzzle = [s.strip() for s in f.readlines()]

def alpha(s):
	"""returns s with letters in alpha order"""
	return "".join(sorted(s))

def normalize(signals):
	for i,s in enumerate(signals):
		signals[i] = alpha(s)
	signals.sort()
	signals.sort(key=len)

def contains(s, needle):
	for letter in needle:
		if not letter in s:
			return False
	return True


def deduce(signals):
	# ['be', 'bde', 'bceg', 'abcdf', 'bcdef', 'cdefg', 'abdefg', 'acdefg', 'bcdefg', 'abcdefg']
	#   0      1      2         3       4        5           6      7         8          9
	one = signals[0]
	four = signals[2]
	seven = signals[1]
	eight = signals[9]
	mapping = {
		one: '1',
		seven: '7',
		four: '4',
		eight: '8',
	}
	# get the six-digit ones
	nine = ''
	for signal in signals:
		if len(signal) == 6:
			#   9 contains all of 4
			if contains(signal, four):
				mapping[signal] = '9'
				nine = signal
			#   0 contains all of 1 and 7
			elif contains(signal, one) and contains(signal, seven):
				mapping[signal] = '0'
			#   6 otherwise
			else:
				mapping[signal] = '6'
	
	# then the five-digit ones
	for signal in signals:
		if len(signal) == 5:
			#   3 contains 7
			if contains(signal, seven):
				mapping[signal] = '3'
			#   2 - 9 doesn't contain it
			elif not contains(nine, signal):
				mapping[signal] = '2'
			else:
				mapping[signal] = '5'
	return mapping


def decode(outputs, mapping):
	result = ""
	for o in outputs:
		result += mapping[alpha(o)]
	return int(result)

total = 0
for line in puzzle:
	(front, back)  = line.split("|")
	signals = [s.strip() for s in front.split()]
	outputs = [s.strip() for s in back.split()]
	normalize(signals)
	mapping = deduce(signals)
	value = decode(outputs, mapping)
	total += value

print(total)
