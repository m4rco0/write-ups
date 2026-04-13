import sys
chars = ""
from fileinput import input
for line in input():
  chars += line

lookup1 = "\n \"#()*+/1:=[]abcdefghijklmnopqrstuvwxyz"
lookup2 = "ABCDEFGHIJKLMNOPQRSTabcdefghijklmnopqrst\n"

out = ""

prev = 0
for char in chars:
  cur = lookup2.index(char.lower())
  out += lookup1[(prev - cur) % 40]
  prev = cur

sys.stdout.write(out)
