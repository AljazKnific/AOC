from numpy import False_, True_

f = open("Input.txt", "r")

class Coordinates:
    def __init__(self,x,y) -> None:
        self.x = x
        self.y = y

HashSet = []
x = 0
x2 = 0
y = 0
y2 = 0
beg = Coordinates(x,y)
HashSet.append(beg)
bool = False

while 1:
    bool = False
    char = f.read(1)
    char2 = f.read(1)
    
    if not char or not char2:
        break
    
   # if x == 0 and char == '<':
    #    print("X cannot be negative value")
     #   break
    #if y == 0 and char == 'v':
    #    print("Y cannot be negative value")
    #    break
    
    if char == 'v':
        y += 1
    elif char == '<':
        x += 1
    elif char == '>':
        x -= 1
    elif char == '^':
        y -= 1
    else:
        print("X")
    
    if char2 == 'v':
        y2 += 1
    elif char2 == '<':
        x2 += 1
    elif char2 == '>':
        x2 -= 1
    elif char2 == '^':
        y2 -= 1
        
    newCoor = Coordinates(x,y)
    newCoor2 = Coordinates(x2,y2)
    for item in HashSet:
        if item.x == x and item.y == y:
            bool = True
            break
    if not bool:
        HashSet.append(newCoor)
    bool = False
    for item in HashSet:
        if item.x == x2 and item.y == y2:
            bool = True
            break
    if not bool:
        HashSet.append(newCoor2)

print(len(HashSet))
    
    