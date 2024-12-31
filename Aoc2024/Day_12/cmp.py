import fileinput

file1 = "output1.txt"
file2 = "output2.txt"

counter = 0

for line1, line2 in zip(fileinput.input(file1), fileinput.input(file2)):

    if counter == 1 and line1 != line2:
        print(f"File1: {line1.strip()}\nFile2: {line2.strip()} (Line: {counter})\n")


    counter += 1  # Increment the counter for each line
    counter %= 2

