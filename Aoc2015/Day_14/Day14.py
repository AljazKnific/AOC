time = 2503

reindeers = []

def get_distance(reindeer, time):
    speed = reindeer["speed"]
    duration = reindeer["duration"]
    rest = reindeer["rest"]

    distance = (speed * duration) * (time // (duration + rest)) + min(time % (duration + rest), duration) * speed
    return distance

with open("Input.txt") as file:
    max_distance = 0
    for line in file:
        line = line.strip().split()
        speed = int(line[3])
        duration = int(line[6])
        rest = int(line[-2])

        reindeers.append({
            "name": line[0],
            "speed": speed,
            "duration": duration,
            "rest": rest,
            "points": 0
        })

        distance = get_distance(reindeers[-1], time)

        if distance > max_distance:
            max_distance = distance


print(max_distance)
#Part two
max_distance = 0
for i in range(time):
    for reindeer in reindeers:
        distance = get_distance(reindeer, i + 1)
        if distance > max_distance:
            max_distance = distance

    for reindeer in reindeers:
        distance = get_distance(reindeer, i + 1)
        if distance == max_distance:
            reindeer["points"] += 1
max_points = 0
for reindeer in reindeers:
    if reindeer["points"] > max_points:
        max_points = reindeer["points"]

print(max_points)