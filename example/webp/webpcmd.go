package main

import (
	"github.com/masayukioguni/go-webp-server/webp"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join("../test-fixtures", "lena.jpg")

	f, _ := os.Open(path)
	defer f.Close()

	m, _ := webp.Decode(f)

	toimg, _ := os.Create("new.webp")
	defer toimg.Close()

	_ = webp.Encode(toimg, m, &webp.Options{false, 50})

}
