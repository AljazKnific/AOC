f = open("Input.txt", "r")

class Pair:
    def __init__(self,x,y) -> None:
        self.name = x
        self.value = y

def getValue(x, set) -> int:
    if isinstance(x, int):
        return x
    elif isinstance(x, str):
        try:
            a = int(x)
            return a
        except ValueError:
            for y in set:
                if y.name == x:
                    return y.value
            return None

def notOper(x) -> int:
    return 65535 - x

def andOper(x,y) -> int:
    x &= y
    return x

def orOper(x,y) -> int:
    x |= y
    return x

def rshiftOper(x, y) -> int:
    x >>= y
    return x

def lshiftOper(x, y) -> int:
    x <<= y
    return x

def getPair(s, name) -> Pair:
    
    for x in s:
        if name == x.name:
            return x

def first():
    
    resSet = set()
    inputList = []
    
    while 1:
        line = f.readline()
        
        if not line:
            break
        
        inst = line.split(' ')
        inst[len(inst) - 1] = inst[len(inst) - 1].replace('\n', '')
        
        if len(inst) == 3:
            if isinstance(inst[0], int):
                resSet.add(Pair(inst[len(inst) - 1], inst[0]))
                
            elif isinstance(inst[0], str):
                try:
                    resSet.add(Pair(inst[len(inst) - 1], int(inst[0])))
                except ValueError:
                    inputList.append(inst)
        else:
            inputList.append(inst)
    
    check = False

    while not check:

        check = True
        for inst in inputList:
            exists = True
            addToList = True
            
            elem1Value = None
            elem2Value = None
            
            if len(inst) == 3:
                elem1Value = getValue(inst[0], resSet)
            elif len(inst) == 4:
                elem1Value = getValue(inst[1], resSet) 
            else:
                if inst[1] == 'AND' or inst[1] == 'OR':
                    elem1Value = getValue(inst[0], resSet)
                    elem2Value = getValue(inst[2], resSet)
                elif inst[1] == 'LSHIFT' or inst[1] == 'RSHIFT':
                    elem1Value = getValue(inst[0], resSet)
                     
            
            elem = getPair(resSet, inst[len(inst) - 1])
        
            if not elem:
                exists = False
            
            addEleM = Pair(inst[len(inst) - 1], 0)
            
            if 'NOT' == inst[0] and elem1Value != None:
                addEleM.value = notOper(elem1Value)
            elif 'AND' == inst[1] and elem1Value != None and elem2Value != None:
                addEleM.value = andOper(elem1Value, elem2Value)
            elif 'OR' == inst[1] and elem1Value != None and elem2Value != None:
                addEleM.value = orOper(elem1Value, elem2Value)
            elif 'LSHIFT' == inst[1] and elem1Value != None:
                addEleM.value = lshiftOper(elem1Value, int(inst[2]))
            elif 'RSHIFT' == inst[1] and elem1Value != None:
                addEleM.value = rshiftOper(elem1Value, int(inst[2]))
            elif '->' == inst[1] and elem1Value != None:
                addEleM.value = elem1Value
            else:
                check = False
                addToList = False
            
            
            if addToList:
                if exists:
                    resSet.remove(elem)
                resSet.add(addEleM)
        
        
    for x in resSet:
        if x.name == 'a':
            print(x.value)
            
first()