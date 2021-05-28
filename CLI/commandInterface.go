package CLI

import (
	ec "cc/errCatch"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func Connect(host RemoteUser) *ssh.Session {
	config := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(host.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//ssh.FixedHostKey(hostKey),
	}

	addr := fmt.Sprintf("%s:%d", host.Host, host.Port)
	client, err := ssh.Dial("tcp", addr, config)
	ec.Out(err, "Ошибка при звонке к SSH")

	session, err := client.NewSession()
	ec.Out(err, "Ошибка при открытии сесиии")

	return session
}
