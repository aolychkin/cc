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

var hostUser = cli.RemoteUser{User: "root", Password: "Xz4lbm777!", Host: "45.90.33.19", Port: 22}

func CoreInstallHandler(w http.ResponseWriter, r *http.Request) {
	client := cli.ServConnect(hostUser)
	var arr []cli.CommandLine
	//arr = append(arr, cli.InstallPreparation)
	//arr = append(arr, cli.InstallPodman)
	//arr = append(arr, cli.InstallNginx)
	//arr = append(arr, cli.InstallSnap)
	arr = append(arr, cli.InstallCertbot)
	RunCommand(client, arr)
}

func main() {
	//http.HandleFunc("/cli", cliHandler)
	http.HandleFunc("/install", CoreInstallHandler)
	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RunCommand(client *ssh.Client, command []cli.CommandLine) {
	go func() {
		for _, x := range command {
			if x.Line == "echo \"ccst\"" {
				break
			}
			session := cli.OpenSession(client)

			finished := make(chan bool)
			fmt.Println("\n[" + client.User() + "] running: " + x.Line)
			go ReadOutputs(session, x.Mission, finished)

			if err := session.Run(x.Line); err != nil {
				log.Print("\n[!ERROR!] Команда выполнилась с ошибкой: " + err.Error())
				//fmt.Printf("%s", str)
			}

			<-finished
			session.Close()
		}
	}()
}

//Чтение вывода консоли в режиме онлайн
func ReadOutputs(session *ssh.Session, mission string, finished chan bool) {
	out, err := session.StdoutPipe()
	ec.Out(err, "\n[!ERROR!] Ошибка при открытии потока вывода")

	r := bufio.NewReader(out)
	var output []byte
	output = []byte(fmt.Sprintln("(-- process --) " + mission + "..."))
	output = append(output, []byte(fmt.Sprintln("<== OUTPUT ==> "))...)
	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			fmt.Printf("%s", output)
			output = nil
		}
		output = append(output, b)
	}
	finished <- true
}
