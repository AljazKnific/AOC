f = open("i.txt", "r")
from collections import defaultdict
dirs = defaultdict(int)
list = []
dirs["/"] = 0
for x in f:
    x = x.strip().split(" ")

    if(x[0] == '$'  and x[1] == 'ls' or x[0] == 'dir'):
        continue

    if(x[0] == '$' and x[1] == 'cd' and x[2] == '..'):
        list.pop()
        continue
    elif(x[0] == '$' and x[1] == 'cd' and x[2] == '/'): 
        list.clear()
        list.append("/")
        continue
    elif(x[0] == '$' and x[1] == 'cd'):
        path = f"{list[-1]}_{x[2]}" if list else x[2]
        list.append(path)
        continue

    for y in list:
        if dirs.get(y) is not None:
            temp = dirs[y] + int(x[0])
            dirs[y] = temp
        else:
            dirs[y] = int(x[0])

regular_dict = dict(dirs)
sorted_dict_by_values = dict(sorted(regular_dict.items(), key=lambda x: x[1]))

for j in sorted_dict_by_values.values():
    if(j > (30000000 - (70000000 - regular_dict["/"]))):
        print(j)
        break
print(sum(value for value in dirs.values() if value <= 100000))
        