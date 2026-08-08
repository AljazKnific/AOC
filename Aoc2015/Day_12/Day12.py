import json

def rec1(data):
    if isinstance(data, int):
        return data
    elif isinstance(data, list):
        return sum(rec1(i) for i in data)
    elif isinstance(data, dict):
        return sum(rec1(i) for i in data.values())

    return 0

def rec2(data):
    if isinstance(data, int):
        return data
    elif isinstance(data, list):
        return sum(rec2(i) for i in data)
    elif isinstance(data, dict):
        if "red" in data.values():
            return 0
        return sum(rec2(i) for i in data.values())

    return 0

with open("Input.txt") as file:
    data = json.load(file)
    print(rec1(data))
    print(rec2(data))


