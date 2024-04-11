f = open("i.txt" , "r")

X = 1
cycle = 1

pairs = {
    20: True,
    60: True, 
    100: True,
    140: True,
    180: True,
    220: True
}

def check(pairs, cycle, X) -> int:
    sum = 0

    if(pairs.get(20) and cycle >= 20):
        sum += X * 20
        pairs[20] = False
    elif(pairs.get(60) and cycle >= 60):
        sum += X * 60
        pairs[60] = False
    elif(pairs.get(100) and cycle >= 100):
        sum += X * 100
        pairs[100] = False
    elif(pairs.get(140) and cycle >= 140):
        sum += X * 140
        pairs[140] = False
    elif(pairs.get(180) and cycle >= 180):
        sum += X * 180
        pairs[180] = False
    elif(pairs.get(220) and cycle >= 220):
        sum += X * 220
        pairs[220] = False

    return sum

def partTwo(cycle, X):
    cycle -= 1
    cycle = cycle % 40
#    print(str(cycle) + " vrednost " + str(X))

    if(cycle >= X and cycle <= X + 2):
        print("#", end="")
    else: 
        print(".", end="")

    if(cycle % 40 == 0):
        print()
        
sum = 0
for x in f:
    x = x.strip()

    if(x == 'noop'):
        #partTwo(cycle, X)
        cycle += 1
        partTwo(cycle, X)
        sum += check(pairs, cycle, X)
    else:
        x = x.split()
        #partTwo(cycle, X)
        cycle += 1
        partTwo(cycle, X)
        sum += check(pairs, cycle, X)
        cycle += 1
        partTwo(cycle, X)
        X += int(x[1])
        sum += check(pairs, cycle, X)
            
        #print(int(x[1]))
print()
print(sum)
