import numpy as num
f = open("i.txt", "r")

grid = []
sum = 0
for x in f: 
    x = x.strip()
    trees = [int(y) for y in x]
    grid.append(trees)

grid = num.array(grid)

#add round elements
bestRes = 0
sum += ((len(grid) - 2) * 2) + (len(grid) * 2)
for i in range(1, len(grid) - 1):
    for j in range(1, len(grid) - 1):
        #print(grid[i][j])
        #levo

        left = max(grid[i][:j])
        right = max(grid[i][j+1:])
        #print(str(i) + " " + str(j))
        if(grid[i][j] > left):
            left = j
        else:
            #print(max([k for k, l in enumerate(grid[i][:j]) if l >= grid[i][j]]))
            left = j - max([k for k, l in enumerate(grid[i][:j]) if l >= grid[i][j]])

        if(grid[i][j] > right):
            right = len(grid) - 1 - j
        else:
            right = min([k for k, l in enumerate(grid[i][j+1:]) if l >= grid[i][j]]) + 1

        #CHECK IF THE VAL IS BIGGer ELSE 

        grid = grid.T
        up = max(grid[j][:i])
        down = max(grid[j][i+1:])

        if(grid[j][i] > up):
            up = i
        else:
            up = i - max([k for k, l in enumerate(grid[j][:i]) if l >= grid[j][i]])

        if(grid[j][i] > down):
            down = len(grid) - 1 - i
        else:
            down = min([k for k, l in enumerate(grid[j][i+1:]) if l >= grid[j][i]]) + 1
        grid = grid.T

        #print("LEFT " + str(left))
        #print("right " + str(right))
        #print("up " + str(up))
        #print("down " + str(down))

        if((left * right * up * down) > bestRes):
            bestRes = left * right * up * down

        if(grid[i][j] > max(grid[i][:j]) or grid[i][j] > max(grid[i][j+1:])):
            sum +=1
            continue

        grid = grid.T

        if(grid[j][i] > max(grid[j][:i]) or grid[j][i] > max(grid[j][i+1:])):
            sum +=1
            grid = grid.T
            continue

        grid = grid.T
print(sum)
print(bestRes)
            