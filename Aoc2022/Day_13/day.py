import ast
f = open("i.txt", "r")

x = list(map(str.splitlines, f.read().strip().split("\n\n")))
print(x)

def rec(x,y):

    if isinstance(x, int):
        if isinstance(y, int):
            return x - y
        else:
            return rec([x], y)
    else:
        if isinstance(y, int):
            return rec(x, [y])

    for a, b in zip(x,y):
        v = rec(a,b)

        if v:
            return v
        
    return len(x) - len(y)

sum = 0

for i ,(l, r) in enumerate(x):
    if rec(eval(l), eval(r)) < 0:
        sum += i + 1

print(sum)

def extract_digits_per_line(x):
    result = []
    for inner_list in x:
        line_result = []
        for inner_str in inner_list:
            try:
                # Safely evaluate the string as a literal expression
                line_result.extend([int(digit) for digit in ast.literal_eval(inner_str) if isinstance(digit, int)])
            except (SyntaxError, ValueError):
                print(f"Error evaluating: {inner_str}")
        result.append(line_result)
    return result

# Extract digits per line from the x
digits_per_line = extract_digits_per_line(x)

print(digits_per_line)
