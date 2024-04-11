f = open("Input.txt", "r")

class Path:
    def __init__(self,x,y) -> None:
        self.name = x
        self.value = y

class City: 
    def __init__(self, name) -> None:
        self.name = name
        self.optimal = 10000000
        self.paths = []
        
    def addPath(self, name, value) -> None:
        if value < self.optimal:
            self.optimal = value
            self.paths.insert(0, Path(name, value))
        else:
            self.paths.append(Path(name, value))
        
def checkCity(name, graph):
        
        for x in graph:
            if x.name == name:
                return x, True
            
        return City(name), False
    
def graphTraversal(graph, name, visited):
    if len(graph) == len(visited):
        return 0
    
    if name == '':
        name = list(graph)[0].name
        visited.add(name)
    
    for x in graph:
        if x.name == name:
            for y in x.paths:
                if y.name not in visited:
                    visited.add(y.name)
                    return y.value + graphTraversal(graph, y.name, visited)
                    break
             
    return 0    
    
def first():
    
    graph = set()
    
    while 1:
        
        line = f.readline()
        
        if not line:
            break
        
        data = line.split(' ')
        data[len(data) - 1] = data[len(data) - 1].replace('\n', '')

        distance = int(data[4])
        cityA, cityACheck = checkCity(data[0], graph)
        cityB, cityBCheck = checkCity(data[2], graph)
        
        cityA.addPath(data[2],distance)
        cityB.addPath(data[0], distance)
        if not cityACheck:
            graph.add(cityA)
        if not cityBCheck:
            graph.add(cityB)
    
    print(graphTraversal(graph, '', set()))    
    for x in graph:
        print(x.name)
        for y in x.paths:
            print("value: " + str(y.value) + " -> " + y.name)
        
    return 0

first()