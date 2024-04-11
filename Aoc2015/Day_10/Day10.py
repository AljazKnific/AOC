f = open("Input.txt", "r")

line = f.readline()

i = 0
while i < 42:
    newString = ""
    x = 0
    while x < len(line):
        num = line[x]
        times = 0

        while 1:
            if x < len(line) and num == line[x]:
                times += 1
                x += 1
            else:
                break
        
        newString += str(times) + str(num)
    line = newString
    i += 1

print(len(newString))
