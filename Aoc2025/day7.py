from functools import cache
f = open("Input.txt", "r")

room = []
for line in f:
    room.append(list(line[:len(line)-1]))

searchList = [(0, x) for x in range(len(room[0])) if room[0][x] == 'S']
searchSet = set(searchList)

def firstExec(searchSet):
    res1 = 0

    while len(searchSet) > 0:
        nextIteration = set()
        for x,y in searchSet:
            
            #Check edges
            if x + 1 >= len(room) or y < 0 or y >= len(room[0]):
                continue
            
            if room[x + 1][y] == "^":
                nextIteration.add((x + 1, y - 1))
                nextIteration.add((x + 1, y + 1))
                res1 += 1
            
            if room[x + 1][y] == ".":
                nextIteration.add((x + 1, y))

        searchSet = nextIteration

    return res1

@cache
def second(x,y):
    if x >= len(room):
        return 1
    
    if room[x][y] == "^":
        return second(x + 1, y - 1) + second(x + 1, y + 1)
    elif room[x][y] == "." or room[x][y] == "S":
        return second(x + 1, y)
    


print(f"Result: {firstExec(searchSet)}")
print(f"Result2: {second(searchList[0][0], searchList[0][1])}")

f.close()


