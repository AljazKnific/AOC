f = open("i.txt", "r")

for x in f:
    index = 0
    for y in x:
        if(index > 13):
            if(len(set(x[index-14:index])) == 14):
                print(index)
                break
        index += 1