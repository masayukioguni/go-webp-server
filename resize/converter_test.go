package resize

import (
	"github.com/masayukioguni/go-webp-server/webp"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReseize_Resize(t *testing.T) {
	testPath := filepath.Join("../test-fixtures", "lena.jpg")
	defer func() {
		_ = os.Remove("resize.webp")
	}()
	f, err := os.Open(testPath)
	defer f.Close()

	m, _ := Resize(f, 2000, 0)

	toimg, _ := os.Create("resize.webp")
	defer toimg.Close()

	err = webp.Encode(toimg, m, &webp.Options{false, 50})

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("TestReseize_Resize  returned %+v", err)
	}
}
