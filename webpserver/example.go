package main

import (
	"github.com/masayukioguni/go-webp-server/server"
)

func main() {
	server := server.NewServer("http://lohas.nicoseiga.jp/")

	server.Run()
}
