cipher_txt = "cvpbPGS{abg_gbb_onq_bs_n_ceboyrz}"

tabela1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
tabela2 = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

saida = "" 
for char in cipher_txt:
    if char not in tabela1:
        saida += char
        continue 
    indice = tabela1.index(char)
    saida += tabela2[indice]
print(saida)