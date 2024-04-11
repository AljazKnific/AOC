f = open("Input.txt", "r")

sum = 0
ribbon = 0

while 1:
    line = f.readline()
    
    if not line:
        break
    
    dim = line.split('x')
    l = int(dim[0])
    w = int(dim[1])
    h = int(dim[2])
    
    x = l * w
    y = w * h
    z = h * l
    smallest = min(x,y,z)
    smallest2 = min(l,w,h)
    mid = min(max(w,l), max(w,h), max(h,l))
    sum += 2 * (x + y + z) + smallest
    
    ribbon += 2 * (smallest2 + mid) + l * w * h
    
print(sum)
print(ribbon)