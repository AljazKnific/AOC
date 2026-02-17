f = open("Input.txt", "r")


def nearest_neighbor_tsp(distances):
    bestDistance = float('inf')

    for start in range(len(distances)):

        n = len(distances)
        visited = [False] * n
        route = [start]
        visited[start] = True
        total_distance = 0

        for _ in range(1, n):
            last = route[-1]
            nearest = None
            min_dist = float('inf')
            for i in range(n):
                if not visited[i] and distances[last][i] < min_dist:
                    min_dist = distances[last][i]
                    nearest = i
            route.append(nearest)
            visited[nearest] = True
            total_distance += min_dist

        if total_distance < bestDistance:
            bestDistance = total_distance

    return route, bestDistance

def most_farthest_neighbor_tsp(distances):
    bestDistance = 0

    for start in range(len(distances)):

        n = len(distances)
        visited = [False] * n
        route = [start]
        visited[start] = True
        total_distance = 0

        for _ in range(1, n):
            last = route[-1]
            farthest = None
            max_dist = float('-inf')
            for i in range(n):
                if not visited[i] and distances[last][i] > max_dist:
                    max_dist = distances[last][i]
                    farthest = i
            route.append(farthest)
            visited[farthest] = True
            total_distance += max_dist

        if total_distance > bestDistance:
            bestDistance = total_distance

    return route, bestDistance


size = 8
matrix = [[0 for i in range(size)] for j in range(size)]

country = 0
currCountry  = ""
index = 1

for line in f:
    line = line.strip().split()
    distance = int(line[4])

    if currCountry == "":
        currCountry = line[0]
    elif currCountry != line[0]:
        country += 1
        currCountry = line[0]
        index = country + 1

    matrix[country][index] = distance
    matrix[index][country] = distance
    index += 1

print(matrix)

route, total_distance = nearest_neighbor_tsp(matrix)
route2, total_distance2 = most_farthest_neighbor_tsp(matrix)
print("Route:", route)
print("Total Distance:", total_distance)

print("Route:", route2)
print("Total Distance:", total_distance2)