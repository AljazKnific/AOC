import hashlib

i = 0
while True:
    s = f"ckczppom{i}"
    h = hashlib.md5(s.encode()).hexdigest()
    if h.startswith("000000"):
        print(s, h)
        break
    i += 1


