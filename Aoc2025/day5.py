f = open("input.txt", "r")

def isContained(x, low, top):
    return x >= low and x <= top

def lowerValue(x, y):
    if x < y:
        return x
    return y

def upperValue(x, y):
    if x > y:
        return x
    return y

ingiridents = False
freshingridients = []
freshingridients2 = []
res1 = 0
res2 = 0
for line in f:
    if line == "\n":
        ingiridents = True
        continue
    if not ingiridents:
        x, y = map(int, line[:len(line)-1].split("-"))
        freshingridients.append((x, y))

        i = 0
        while i  < len(freshingridients2):
            low, top = freshingridients2[i]


            is_overlapping = (isContained(x, low, top) or isContained(y, low, top) or isContained(low, x, y) or isContained(top, x, y))
            if is_overlapping:
                hitItem = freshingridients2.pop(i)
                i = 0
                x = lowerValue(x, hitItem[0])
                y = upperValue(y, hitItem[1])
            else:
                i += 1

        freshingridients2.append((x, y))

    else:
        x = int(line[:len(line)-1])
        for low,top in freshingridients:
            if x >= low and x <= top:
                res1 += 1
                break

print(f"Result: {res1}")
for x,y in freshingridients2:
    res2 += (y - x + 1)
print(f"Result2: {res2}")
f.close()