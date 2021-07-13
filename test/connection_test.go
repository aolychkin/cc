package test

import (
	cli "cc/CLI"
	"testing"
)

var hostUser = cli.RemoteUser{User: "root", Password: "Xz4lbm777!", Host: "45.90.33.19", Port: 22}

func BenchmarkConnect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//478771120 ns/op     5105 B/op         445 allocs/op 	Session
		//239151875 ns/op     36004 B/op        345 allocs/op 	Serv
		//125719350 ns/op     11882 B/op        121 allocs/op	Open Session
		cli.ServConnect(hostUser)
	}
}

func BenchmarkOpenSession(b *testing.B) {
	client := cli.ServConnect(hostUser)
	for i := 0; i < b.N; i++ {
		cli.OpenSession(client)
	}
}