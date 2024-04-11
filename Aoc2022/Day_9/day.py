import numpy as np
f = open("i.txt" , "r")

def calculate(sez, coords, x, y):

    sez[0] = (sez[0][0] + x, sez[0][1] + y)

    for j in range(0,9):
        H = sez[j]
        T = sez[j + 1]

        xD = H[0] - T[0]
        yD = H[1] - T[1]

        if(abs(xD) > 1 or abs(yD) > 1):
            if(xD == 0):
                T[1] += yD // 2
            elif(yD == 0):
                T[0] += xD // 2
            else:
                T[0] += 1 if xD > 0 else -1
                T[1] += 1 if yD > 0 else -1

    coords.add(tuple(sez[-1]))
    return

sez = [[0,0] for _ in range(0,10)]
H = [0,0]
T = [0,0]

coords = set()
coords.add(tuple(T))
for x in f:
    x = x.strip().split()
    for i in range(0, int(x[1])):
        prevH = H
        if(x[0] == 'R'):
            calculate(sez, coords, 0, 1)
            #H = (H[0], H[1] + 1)
        elif(x[0]  == 'L'):
            calculate(sez, coords, 0, -1)
            #H = (H[0], H[1] - 1)
        elif(x[0]  == 'U'):
            calculate(sez, coords, -1, 0)
            
            #H = (H[0] + 1, H[1])
            
        elif(x[0]  == 'D'):
            calculate(sez, coords, 1, 0)
            #H = (H[0] - 1, H[1])

        #if(abs(T[0] - H[0]) >= 2 or  abs(T[1] - H[1]) >= 2):
           # T = prevH
           # coords.add(T)
#print(coords)
print(len(coords))