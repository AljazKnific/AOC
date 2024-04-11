f = open("i.txt", "r")

#we create list 24 lists long
list = [[] for _ in range(3)]
check = True

for x in f:
    if(len(x) == 1 or x[1] == '1'):
        check = False
        continue

    if(check):
        j = 1
        index = 0
        while j < len(x):
            if(x[j] != ' '):
                list[index].append(x[j])
            j += 4
            index += 1
    else:
        inst = x.strip().split(" ")
        for q in range(int(inst[1])):
            temp = list[int(inst[3]) - 1].pop(0)
            list[int(inst[5]) - 1].insert(q,temp)

print(*(list[x][0] for x in range(len(list))))


