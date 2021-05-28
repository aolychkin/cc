package main

import (
	"bufio"
	cli "cc/CLI"
	ec "cc/errCatch"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
)
//go build -o a.exe && a.exe
var hostUser = cli.RemoteUser{User: "root", Password: "Xz4lbm777!", Host: "45.90.33.19", Port: 22}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	RunCommand(hostUser, "ls", "Show files")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Запуск списка команд
func RunCommand(host cli.RemoteUser, command, mission string) {
	session := cli.Connect(host)
	defer session.Close()

	finished := make(chan bool)
	fmt.Println("[" + host.User + "] running: " + command)
	go outputInProcess(session, mission, finished)

	if err := session.Run(command); err != nil {
		log.Print("[!ERROR!] Failed to run: " + err.Error())
	}
	<-finished
}

//Чтение вывода консоли в режиме онлайн
func outputInProcess(session *ssh.Session, mission string, finished chan bool) {
	out, err := session.StdoutPipe()
	ec.Out(err, "\n[!ERROR!] Ошибка при открытии потока вывода")

	r := bufio.NewReader(out)
	var output []byte
	fmt.Println("[process] " + mission + "...")
	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		output = append(output, b)

		if b == byte('\n') {
			fmt.Print(string(output))
			output = nil
			continue
		}
	}
	finished <- true
}
