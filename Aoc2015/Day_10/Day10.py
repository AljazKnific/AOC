f = open("Input.txt", "r")

line = f.readline()

i = 0
while i < 40:
    parts = []
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
        
        parts.append(str(times))
        parts.append(num)
    line = "".join(parts)
    i += 1

print(len(line))
