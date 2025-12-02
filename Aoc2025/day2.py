f = open("input.txt", "r")

res1 = 0
res2 = 0
for line in f:
    
    parts = line.replace("\n", "").split(',')
    for part in parts:
        if part == "":
            continue
        x, y = part.split('-')
        
        #Part 1
        for i in range(int(x), int(y)+1):
            temp = str(i)
            if len(temp) % 2 == 1:
                continue
            if temp[:len(temp)//2] == temp[len(temp)//2:]:
                res1 += int(temp)

        #Part 2
        for i in range(int(x), int(y)+1):
            temp = str(i)
            
            for k in range(1, len(temp)):
                bool = False
                if len(temp) % k == 0:
                    bool = True
                    for j in range(1, len(temp)//k):
                        bool &= temp[(j-1)*k:j*k] == temp[j*k:(j+1)*k]
                if bool:
                    res2 += int(temp)
                    break
                

print(f"Result: {res1}")
print(f"Result: {res2}")
f.close()