package main

import (
	cli "cc/CLI"
	"testing"
)

// RunCommand 				167418100 ns/op            24300 B/op        203 allocs/op
// RunCommandConsole 		 62961650 ns/op            10290 B/op        98 allocs/op
// RunCommandConsoleOnce 	 59701675 ns/op            9958 B/op         89 allocs/op
func BenchmarkRunCommand(b *testing.B) {
	client := cli.ServConnect(hostUser)
	var arr []cli.CommandLine
	arr = append(arr, cli.Ls)
	arr = append(arr, cli.Date)
	for i := 0; i < b.N; i++ {
		RunCommand(client, arr)
	}
}
