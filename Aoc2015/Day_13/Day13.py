import json
from itertools import permutations

people = ["Alice", "Bob", "Carol", "David", "Eric", "Frank", "George", "Mallory", "Alyo"]

happiness = dict()

with open("Input.txt") as file:
    for line in file:
        line = line.strip().split()

        happy = int(line[3])
        if line[2] == "lose":
            happy *= -1

        happiness[(line[0], line[-1][:-1])] = happy
max_happiness = float("-inf")

for i in range(len(people) - 1):
    happiness[(people[i], "Alyo")] = 0
    happiness[("Alyo", people[i])] = 0


for arrangement in permutations(people[1:]):
    arrangement = ("Alice",) + arrangement
    total_happiness = 0
    for i in range(len(arrangement)):
        total_happiness += happiness[(arrangement[i], arrangement[(i + 1) % len(arrangement)])]
        total_happiness += happiness[(arrangement[(i + 1) % len(arrangement)], arrangement[i])]
        
    if total_happiness > max_happiness:
        max_happiness = total_happiness

print(max_happiness)