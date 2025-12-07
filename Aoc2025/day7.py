from dataclasses import dataclass
f = open("Input.txt", "r")

@dataclass(frozen=True)
class Segment:
    lastHit: tuple
    currPosition: tuple


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


def secondExec(searchSet):
    res2 = 0
    while len(searchSet) > 0:
        #print(searchSet)
        #print(f"Segments: {len(searchSet)}")
        nextIteration = set()
        for segment in searchSet:
            (x,y) = segment.currPosition
            (lastX, lastY) = segment.lastHit
            
            #Check edges
            if x + 1 >= len(room) or y < 0 or y >= len(room[0]):
                continue
            
            if room[x + 1][y] == "^":
                nextIteration.add(Segment((x+1,y), (x + 1, y - 1)))
                nextIteration.add(Segment((x+1,y), (x + 1, y + 1)))
                res2 += 1
            
            if room[x + 1][y] == ".":
                nextIteration.add(Segment((lastX, lastY), (x+1,y)))

        if len(nextIteration) == 0:
            print(len(searchSet))

        searchSet = nextIteration


    return res2

print(f"Result: {firstExec(searchSet)}")
segments = {
    Segment((-1,-1), (0, x))
    for x in range(len(room[0]))
    if room[0][x] == 'S'
}

print(f"Result: {secondExec(segments)}")
f.close()


