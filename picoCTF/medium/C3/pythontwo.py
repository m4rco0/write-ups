import fileinput
chars = ""
for line in fileinput.input():
    chars += line

b = 1 
for i in range(len(chars)):
    if i == pow(b, 3):      # i = b³
        print(chars[i]) #prints
        b += 1 
