f = open("i.txt", "r")
# A - rock B - paper C- sc
# X - rock 1 Y - paper 2 Z - sc 3 
#Loss 0 Draw 3 Win 6
possibleDraws = {
    "A Y" : 8,
    "A X" : 4,
    "A Z" : 3,   
    "B Y" : 5,
    "B X" : 1,
    "B Z" : 9,
    "C Y" : 2,
    "C X" : 7,
    "C Z" : 6
}
#X - lose Y- draw Z - win
possibleLoses = {
    "A X" : 3,
    "A Y" : 3 + 1,
    "A Z" : 6 + 2,   
    "B X" : 0 + 1,
    "B Y" : 3 + 2,
    "B Z" : 6 + 3,
    "C X" : 0 + 2,
    "C Y" : 3 + 3,
    "C Z" : 6 + 1
}


num = 0
num2 = 0
for x in f:
    x = x.strip()
    num += possibleDraws.get(x)
    num2 += possibleLoses.get(x)


print(num)
print(num2)