import numpy as np

f = open("i.txt", "r")
items = []
operations = []
divisible = []
tests = []

index = -1

for x in f:
    if len(x) == 1:
        continue
    x = x.strip()
    x = x.split()

    if x[0] == 'Monkey':
        index += 1
        continue

    if x[0] == 'Starting':
        items.append([np.int64(k.replace(",", "")) for k in x[2:]])

    if x[0] == 'Operation:':
        operations.append([k for k in x[4:]])

    if x[0] == 'Test:':
        divisible.append(np.int64(x[3]))

    if x[1] == 'true:':
        tests.append([int(x[5]), 0])

    if x[1] == 'false:':
        tests[index][1] = int(x[5])

monkeys = [0 for _ in range(index + 1)]
mod = np.prod(divisible)

for m in range(10000):
    for i in range(len(items)):

        length = len(items[i])
        for _ in range(length):
            old = items[i].pop(0)
            new = 0
            if operations[i][1] == "old":
                sec = old
            else:
                sec = int(operations[i][1])

            if operations[i][0] == '*':
                new = (old * sec) % mod

            elif operations[i][0] == '+':
                new = (old + sec) % mod
            if new % divisible[i] == 0:
                items[tests[i][0]].append(new)
            else:
                items[tests[i][1]].append(new)
            monkeys[i] += 1
print(monkeys)
res = max(monkeys)
monkeys.remove(res)
res *= max(monkeys)
print(res)
