
rescontainers = []

def rec(index, remaining, containers): 
    if remaining == 0:  
        rescontainers.append(containers[:])
        return 1

    if remaining < 0 or index == len(water): 
        return 0

    return rec(index + 1, remaining, containers) + rec(index + 1, remaining - water[index], containers + [water[index]])

    

water = [int(line.strip()) for line in open("input.txt")]
volume = 150
print(rec(0, volume, []))

min_len = min(len(x) for x in rescontainers)
answer = sum(len(x) == min_len for x in rescontainers)
print(answer)