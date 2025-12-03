def main():
    f = open("input.txt", "r")
    res1 = 0
    res2 = 0
    for bank in f:
        #Dont get the last \n and last number
        top, pos = getBatteryAndPosition(bank[:len(bank) - 2])
        top2, _ = getBatteryAndPosition(bank[pos+1:len(bank) - 1])

        res1 += top * 10 + top2

        #Part 2
        secPos = 0
        bankRes = 0
        for i in reversed(range(0,12)):
            x, y = getBatteryAndPosition(bank[secPos:len(bank) - 1 - i])
            secPos += y + 1
            bankRes += x * (10 ** i)

        res2 += bankRes
    print(f"Result: {res1}")
    print(f"Result: {res2}")
    f.close()

def getBatteryAndPosition(bank):
    top = -1
    pos = -1

    for iBattery , battery in enumerate(bank):
        if int(battery) > top:
            top = int(battery)
            pos = iBattery
    return top, pos

main()