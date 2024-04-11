import collections
import numpy
f = open("Input.txt", "r")
arr = numpy.zeros((1000, 1000))

def first():
    
    while 1:
        line = f.readline()
        
        if not line:
            break
        
        inst = line.split(' ')
        x = 2
        
        if inst[0] == 'toggle':
            x = 1
        start = inst[x].split(',')
        end = inst[x + 2].split(',')
        
        
        for x in range(int(start[0]), int(end[0]) + 1):
            for y in range(int(start[1]), int(end[1]) + 1):
                if 'on' == inst[1]:
                    arr[x][y] = 1
                elif 'off' == inst[1]:
                    arr[x][y] = 0
                elif 'toggle' == inst[0]:
                    if arr[x][y] == 1:
                        arr[x][y] = 0
                    else:
                        arr[x][y] = 1
        



def second():
    while 1:
        line = f.readline()
        
        if not line:
            break
        
        inst = line.split(' ')
        x = 2
        
        if inst[0] == 'toggle':
            x = 1
        start = inst[x].split(',')
        end = inst[x + 2].split(',')
        
        
        for x in range(int(start[0]), int(end[0]) + 1):
            for y in range(int(start[1]), int(end[1]) + 1):
                if 'on' == inst[1]:
                    arr[x][y] += 1
                elif 'off' == inst[1]:
                    if arr[x][y] > 0:
                        arr[x][y] -= 1
                elif 'toggle' == inst[0]:
                    arr[x][y] += 2
                
        

        
#first()
#print(numpy.count_nonzero(arr == 1))
second()
print(numpy.sum(arr))