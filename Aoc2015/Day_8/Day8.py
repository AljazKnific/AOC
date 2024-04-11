f = open("Input.txt", "r")

# First part
def rec(line) -> int:
    num = 0
    
    i = 1
    while i < len(line) - 1:
        if line[i] == '\\':
            if line[i + 1] == 'x':
                num += 1
                i += 4
            elif line[i + 1] == '\"' or line[i + 1] ==  '\\' :
                num += 1
                i += 2
        else:
            num += 1
            i += 1

    return num

# Second part
def rec2(line) -> str:
    # starting quote
    res = '\"'
    res += '\\\"'
    
    i = 1
    while i < len(line) - 1:
        if line[i] == '\\':
            res+= '\\\\'
            if line[i + 1] == 'x':
                res += 'x'
                i += 2
            elif line[i + 1] == '\"':
                res += '\\\"'
                i += 2
            elif line[i + 1] ==  '\\':
                res += '\\\\'    
                i += 2
        else:
            res+= line[i]
            i += 1
            
    #ending quote
    res += '\\\"'
    res += '\"'
    return res

def first():
    numChars = 0
    numLetters = 0
    numCharsSec = 0
    while 1:
        line = f.readline()
        
        if not line:
            break
        
        line = line[:-1] #remove '\n'
        numLetters += len(line)
        numChars += rec(line)
        numCharsSec += len(rec2(line))
        
    print("Result: " + str(numLetters - numChars))
    print("Result2: " + str(numCharsSec - numLetters))
    return 1

first()