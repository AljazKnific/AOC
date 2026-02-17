input = "hepxcrrq"

map = {
    "a" : 0,
    "b" : 1,
    "c" : 2,
    "d" : 3,
    "e" : 4,
    "f" : 5,
    "g" : 6,
    "h" : 7,
    "i" : 8,
    "j" : 9,
    "k" : 10,
    "l" : 11,
    "m" : 12,
    "n" : 13,
    "o" : 14,
    "p" : 15,
    "q" : 16,
    "r" : 17,
    "s" : 18,
    "t" : 19,
    "u" : 20,
    "v" : 21,
    "w" : 22,
    "x" : 23,
    "y" : 24,
    "z" : 25
}

def firstCriteria(password):
    for i in range(len(password) - 2):
        if password[i] + 1 == password[i + 1] and password[i + 1] + 1 == password[i + 2]:
            return True
    return False
def secondCriteria(password):
    for i in password:
        if i == 14 or i == 8 or i == 11:
            return False
    return True

def thirdCriteria(password):
    pairs = 0
    i = 0
    while i < len(password) - 1:
        if password[i] == password[i + 1]:
            pairs += 1
            i += 1
        i += 1
    return pairs >= 2

def printPassword(password):
    result = ""
    for i in password:
        for key, value in map.items():
            if value == i:
                result += key
    print(result)

def findSolution(password):
    second = False
    while 1:
        for i in range(len(password) -1):
            password[len(password) - 1 - i] = (password[len(password) - 1 - i] + 1) % 26
            if password[len(password) - 1 - i] != 0:
                break

        
        if firstCriteria(password) and secondCriteria(password) and thirdCriteria(password):
            if second:
                return password
            second = True


mainList = []
for c in input:
    mainList.append(map[c])
print(mainList)
solution = findSolution(mainList)
printPassword(solution)