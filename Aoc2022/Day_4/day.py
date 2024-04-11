f = open("i.txt", "r")
sum = 0
sumPartTwo = 0
for x in f:
    x = x.split(",")
    a,b = x[0].split("-")
    c,d = x[1].split("-")
    a = int(a)
    b = int(b)
    c = int(c)
    d = int(d)    

    if(a <= c and d <= b or c <= a and b <= d):
        #print("Smaller " + a + " " + b + " : " + c + " " + d)
        sum += 1
        #continue

    if not(d < a or c > b):
        #print("Smaller " + str(a) + " " + str(b) + " : " + str(c) + " " + str(d))
        
        sumPartTwo += 1

print(sum)
print(sumPartTwo)