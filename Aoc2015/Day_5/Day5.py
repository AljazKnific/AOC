f = open("Input.txt", "r")

def vowels(a) -> bool:
    vow = {'a', 'e', 'i', 'o', 'u'}
    
    return a in vow

def dup(a,b) -> bool:    
    return a == b

def naughty(a,b)->bool:
    if a == 'a' and b == 'b':
        return True
    if a == 'c' and b == 'd':
        return True
    if a == 'p' and b == 'q':
        return True
    if a == 'x' and b == 'y':
        return True
    return False

def partOne(f):
    niceCounter = 0

    while 1:
        line = f.readline()

        if not line:
            break

        vowCond = 0
        dupCond = False
        forr = True
            
        for i in range(0, len(line)):
            
            if vowels(line[i]):
                
                vowCond +=1
            if i != len(line) - 1:
                if dup(line[i], line[i+1]):
                    
                    dupCond = True
                if naughty(line[i], line[i+1]):
                    
                    forr = False
                    
        if vowCond >= 3 and dupCond and forr:
            niceCounter += 1
        
    print(niceCounter)
    
    
def partTwo(f):
    
    nice = 0
    while 1:
        
        line = f.readline()
        
        overlap = 0
        check = False
        wordSet = set()
        lastWord = ''
        
        if not line:
            break
            
        for x in range(0, len(line)):
            
            if x > 1:
                if dup(line[x-2], line[x]):
                    overlap += 1
            
            if x == 1:
                currWord = line[x-1] + line[x] 
                wordSet.add(currWord)
            elif x > 1:
                lastWord = line[x-2] + line[x-1]
                currWord = line[x-1] + line[x]
                
                if currWord in wordSet and lastWord != currWord:
                    check = True
                wordSet.add(currWord)
        
        if check and overlap > 0:
            nice += 1
    print(nice)
                       
    

#partOne(f)
partTwo(f)
        
        