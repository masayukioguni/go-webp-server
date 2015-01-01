package main

import (
	"github.com/masayukioguni/go-webp-server/resize"
	"github.com/masayukioguni/go-webp-server/webp"

	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join("../test-fixtures", "lena.jpg")

	f, _ := os.Open(path)
	defer f.Close()

	m, _ := resize.Resize(f, 2000, 0)

	toimg, _ := os.Create("resize.webp")
	defer toimg.Close()

	_ = webp.Encode(toimg, m, &webp.Options{false, 50})

}
