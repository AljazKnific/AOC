
import math
def listOfShortestDistances(boxes):
    disList = []
    
    n = len(boxes)
    for i in range(n-1):
        for j in range(i+1, n):
            dx = boxes[i][0] - boxes[j][0]
            dy = boxes[i][1] - boxes[j][1]
            dz = boxes[i][2] - boxes[j][2]
            distance_sq = dx*dx + dy*dy + dz*dz
            disList.append([i, j, distance_sq])
    
    disList.sort(key=lambda x: x[2])
    return disList


def findParent(parents, node):
    if parents[node] != node:
        parents[node] = findParent(parents, parents[node])
    return parents[node]

def first():
    parents = [x for x in range(len(boxes))]
    for _ in range(connections):
        edge = distances.pop(0)
        box1 = edge[0]
        box2 = edge[1]

        parent1 = findParent(parents, box1)
        parent2 = findParent(parents, box2)

        if parent1 != parent2:
            parents[parent2] = parent1

    res1 = [0 for _ in range(len(boxes))]

    for i, num in enumerate(parents):
        root = findParent(parents, num)
        res1[root] += 1

    print((math.prod(sorted(res1, reverse=True)[:3])))

def second():
    parents = [x for x in range(len(boxes))]
    circuits = len(boxes)
    while circuits > 1:
        edge = distances.pop(0)
        box1 = edge[0]
        box2 = edge[1]

        parent1 = findParent(parents, box1)
        parent2 = findParent(parents, box2)

        if parent1 != parent2:
            parents[parent2] = parent1
            circuits -= 1

        if circuits == 1:
            print(boxes[box1][0] * boxes[box2][0])

f = open("Input.txt", "r")

boxes = []
connections = 10

for line in f:

    cords = list(map(int, line.strip().split(",")))
    boxes.append(cords)

distances = listOfShortestDistances(boxes)
#first()
second()