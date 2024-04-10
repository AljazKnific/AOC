f = open("i.txt", "r")

sum = 0
index = 0
res = []
for x in f:
    index += 1
    x = x.strip()
    res.append(x)
    #
    #firstpart, secondpart = x[:len(x)//2], x[len(x)//2:]
    #res = set(firstpart).intersection(secondpart)
    #temp = res.pop()
    #
    #   if(temp.islower()):
    #      sum += (ord(temp) - 96)
    #
    #   if(temp.isupper()):
    #      sum += (ord(temp) - 65 + 27)

    if(index == 3):
        index = 0
        res = set(res[0]).intersection(res[1]).intersection(res[2])
        temp = res.pop()
        if(temp.islower()):
          sum += (ord(temp) - 96)
    
        if(temp.isupper()):
          sum += (ord(temp) - 65 + 27)
        res = []


print(sum)
