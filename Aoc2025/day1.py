point = 50
res = 0
cross = 0

f = open("day1_input.txt", "r")

for line in f:
    move = line[0]
    steps = int(line[1:len(line)-1 ])
    cross += steps // 100
    steps = steps % 100

    if move == "L":
        temp = point - steps
        if 0 < point and temp <= 0:
            cross += 1
        point = temp % 100

    elif move == "R":
        temp = point + steps
        if temp >= 100:
            cross += 1
        point = temp % 100
    if point == 0:
        res += 1

print(f"Result: {res}")
print(f"Crossings: {cross}")
f.close()
    
