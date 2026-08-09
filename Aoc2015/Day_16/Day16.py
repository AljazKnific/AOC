rules = {
    "children" : 3,
    "cats" : 7,
    "samoyeds" : 2,
    "pomeranians" : 3,
    "akitas" : 0,
    "vizslas" : 0,
    "goldfish" : 5,
    "trees" : 3,
    "cars" : 2,
    "perfumes" : 1
}

def check_rule1(rule, val):
    return rules[rule] == val

def check_rule2(rule, val):
    if rule in ["cats", "trees"]:
        return rules[rule] < val
    elif rule in ["pomeranians", "goldfish"]:
        return rules[rule] > val
    else:
        return rules[rule] == val

with open("Input.txt") as file:
    for line in file:
        line = line.strip().split()
        num = line[1][:-1]

        if not check_rule1(line[2][:-1], int(line[3][:-1])) or not check_rule1(line[4][:-1], int(line[5][:-1])) or not check_rule1(line[6][:-1], int(line[7])):
            continue
        print(num)
        break

with open("Input.txt") as file:
    for line in file:
        line = line.strip().split()
        num = line[1][:-1]

        if not check_rule2(line[2][:-1], int(line[3][:-1])) or not check_rule2(line[4][:-1], int(line[5][:-1])) or not check_rule2(line[6][:-1], int(line[7])):
            continue
        print(num)
        break