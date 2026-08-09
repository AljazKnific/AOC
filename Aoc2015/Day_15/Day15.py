ingridients = []

with open("Input.txt") as file:
    for line in file:
        line = line.strip().split()
        ingridients.append({
            "name": line[0][:-1],
            "capacity": int(line[2][:-1]),
            "durability": int(line[4][:-1]),
            "flavor": int(line[6][:-1]),
            "texture": int(line[8][:-1]),
            "calories": int(line[10])
        })

best_score1 = 0
best_score2 = 0

for i in range(100):
    for j in range(100 - i):
        for k in range(100 - i - j):
            l = 100 - i - j - k
            capacity = ingridients[0]["capacity"] * i + ingridients[1]["capacity"] * j + ingridients[2]["capacity"] * k + ingridients[3]["capacity"] * l
            durability = ingridients[0]["durability"] * i + ingridients[1]["durability"] * j + ingridients[2]["durability"] * k + ingridients[3]["durability"] * l
            flavor = ingridients[0]["flavor"] * i + ingridients[1]["flavor"] * j + ingridients[2]["flavor"] * k + ingridients[3]["flavor"] * l
            texture = ingridients[0]["texture"] * i + ingridients[1]["texture"] * j + ingridients[2]["texture"] * k + ingridients[3]["texture"] * l
            calories = ingridients[0]["calories"] * i + ingridients[1]["calories"] * j + ingridients[2]["calories"] * k + ingridients[3]["calories"] * l

            curr_score = capacity * durability * flavor * texture
            if capacity < 0 or durability < 0 or flavor < 0 or texture < 0:
                curr_score = 0
            if curr_score > best_score1:
                best_score1 = curr_score
            if curr_score > best_score2 and calories == 500:
                best_score2 = curr_score

print(best_score1)
print(best_score2)