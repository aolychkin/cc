package CLI

import (
	ec "cc/errCatch"
	"fmt"
	"golang.org/x/crypto/ssh"
)

//Подключение к серверу по логам пользователя, обработка ошибок
func ServConnect(host RemoteUser) *ssh.Client {
	config := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(host.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", host.Host, host.Port)
	client, err := ssh.Dial("tcp", addr, config)
	ec.Out(err, "Ошибка при звонке к SSH")

	return client
}

//Открытие новой сессии, обработка ошибок
func OpenSession(client *ssh.Client) *ssh.Session{
	session, err := client.NewSession()
	ec.Out(err, "Ошибка при открытии сесиии")
	return session
}
