import fileinput
chars = ""
for line in fileinput.input():
    chars += line
b = 1 / 1
for i in range(len(chars)):
    if i == b * b * b:
        print(chars[i]) #prints
        b += 1 / 1

