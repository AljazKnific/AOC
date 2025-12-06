
def calcColResult(col, operation):
    result = 1 if operation == "*" else 0
    for num in col:
        if operation == "*":
            result *= num
        elif operation == "+":
            result += num
    return result


def calcResult(array, operationArray):
    result = 0
    
    for i, op in enumerate(operationArray):
        if op == "*":
            result += calcColResult(array[i], "*")
        elif op == "+":
            result += calcColResult(array[i], "+")

    return result

def secondExec(array, operations):
    result = 0
    transposed = [''.join(row).replace(' ', '') for row in zip(*array)]

    index = 0
    for op in operations:

        tempRes = 1 if op == '*' else 0
        while index < len(transposed) and transposed[index] != '':
            if op == '*':
                tempRes *= int(transposed[index])
            else:
                tempRes += int(transposed[index])
            index += 1
        index += 1
        result += tempRes
    
    return result

f = open("input.txt", "r")

numArray = []
numArray2 = []

res1 = 0
res2 = 0

for line in f:
    temp = line[:len(line)-1].split()
    
    if temp[0] == "*" or temp[0] == "+":
       res1 += calcResult(numArray, temp)
       res2 += secondExec(numArray2, temp)
    else:
        numArray2.append(list(line[:len(line)-1]))
        for i, number in enumerate(temp):
            if len(numArray) <= i:
                numArray.append([])
            numArray[i].append(int(number))

print(f"Result: {res1}")
print(f"Result2: {res2}")
f.close()