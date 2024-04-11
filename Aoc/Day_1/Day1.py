f = open("Input.txt", "r")
floor = 0
index = 1

while 1:
    
    char = f.read(1)
    if not char:
        break
    if(char == '('):
        floor += 1
    if(char == ')'):
        floor -= 1
        
    if(floor < 0):
        break
    index += 1
        
print(floor)
print(index)
f.close()