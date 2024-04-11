f = open("input.txt", "r")
num = 0
maxres = []
for x in f:
    if(x == '\n'):
        maxres.append(num)
        num = 0
        continue

    if(x[-1] == '\n'):
        x = x[:-1]
    x = int(x)
    num += x    

print(max(maxres))
print(sum(sorted(maxres, reverse=True)[0:3]))