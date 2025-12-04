def copyArray(arr):
    newArr = []
    for i in range(0, len(arr)):
        row = []
        for j in range(0, len(arr[i])):
            row.append(arr[i][j])
        newArr.append(row)
    return newArr

def printArray(arr):
    for i in range(0, len(arr)):
        row = ""
        for j in range(0, len(arr[i])):
            row += arr[i][j]
        print(row)


f = open("input.txt", "r")

rollsArray = []
neighbours = [(-1, -1), (-1, 0), (-1, 1),
              (0, -1),          (0, 1),
              (1, -1),  (1, 0),  (1, 1)]

for line in f:
    rollsArray.append(line[:len(line) - 1])
rollsArray2 = copyArray(rollsArray)


res1 = 0
run = True
while run:
    iRollsNum = 0
    for i in range(0, len(rollsArray[0])):
        for j in range(0, len(rollsArray[i])):

            if rollsArray[i][j] == "@":
                rollsAround = 0
                for neighbour in neighbours:
                    ni = i + neighbour[0]
                    nj = j + neighbour[1]
                    if ni >= 0 and ni < len(rollsArray) and nj >= 0 and nj < len(rollsArray[ni]):
                        if rollsArray[ni][nj] == "@":
                            rollsAround += 1
                
                if rollsAround < 4:
                    rollsArray2[i][j] = "."
                    iRollsNum += 1

    if iRollsNum == 0:
        run = False
    else:
        res1 += iRollsNum
        rollsArray = copyArray(rollsArray2)


print(f"Result: {res1}")                       


f.close()

