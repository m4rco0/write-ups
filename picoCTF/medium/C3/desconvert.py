lookup1 = "\n \"#()*+/1:=[]abcdefghijklmnopqrstuvwxyz"
lookup2 = "ABCDEFGHIJKLMNOPQRSTabcdefghijklmnopqrst\n"

from fileinput import input
prev = 0
cipher_txt = ""
saida = ""
for line in input():
    cipher_txt += line
for char in cipher_txt:
    indice_saida = lookup2.index(char)
    cur = (indice_saida + prev) % 40
    prev = cur
    saida += lookup1[cur]

print(saida)
