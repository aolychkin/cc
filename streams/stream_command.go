package streams

import (
	"bufio"
	cli "cc/CLI"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

func RunStream(client *ssh.Client, command, mission string) {
	session := cli.OpenSession(client)
	defer session.Close()

	finished := make(chan bool)
	rPipe, wPipe, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	//При выводе чего-то из консоли это записывается в пайп
	session.Stdout = wPipe
	session.Stderr = wPipe

	//ccst - потоковое стоп слово
	if err := session.Run(command + " ; echo ccst"); err != nil {
		log.Print("[!ERROR!] Failed to run: " + err.Error())
	}

	go writeOutput(rPipe, finished)

	<-finished
}

func writeOutput(input io.Reader, finished chan bool) {
	in := bufio.NewScanner(input)
	for in.Scan() {
		line := in.Text()
		if line != "ccst"{
			//fmt.Println(line)
		} else {
			break
		}
	}
	finished <- true
}