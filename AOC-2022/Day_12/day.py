import numpy as np
import heapq
f = open("i.txt", "r")

def checkNeigh(grid, x ,y, c) -> bool:
    if x < 0 or x >= len(grid):
        return False
    if y < 0 or y >= len(grid[0]):
        return False
    
    symbol = grid[x][y]

    if ord(c) - ord(symbol) > 1:
        return False
    
    return True

def res(grid, start) -> int:
    pq = []
    heapq.heappush(pq, (0, (start[0], start[1])))

    already = { (start[0], start[1])}

    while pq:
        d, coord  = heapq.heappop(pq) 
        if grid[coord[0]][coord[1]] == 'a' or grid[coord[0]][coord[1]] == 'S':
            return d
        
        pairs = [(coord[0], coord[1] + 1), (coord[0], coord[1]-1), (coord[0]-1, coord[1]), (coord[0] + 1, coord[1])]
        for x, y in pairs:
            if checkNeigh(grid, x, y, grid[coord[0]][coord[1]]) and (x , y ) not in already:
                heapq.heappush(pq ,(d + 1, (x , y)))
                already.add((x,y))
    
    return 1

grid = []
start = []
index = 0
for x in f:
    x = x.strip()
    temp = []
    for i in range(len(x)):
        if x[i] == 'E':
            start = [index, i]
            temp.append('z')
        else:
            temp.append(x[i])
            
    grid.append(temp)
        
    index += 1

print(res(grid, start))
