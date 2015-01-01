package webp

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWebp_Decode(t *testing.T) {
	testPath := filepath.Join("../test-fixtures", "lena.jpg")

	f, err := os.Open(testPath)
	defer f.Close()

	_, err = Decode(f)

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("TestWebp_Decode returned %+v, want %+v", err, nil)
	}
}

func TestWebp_Encode(t *testing.T) {
	testPath := filepath.Join("../test-fixtures", "lena.jpg")
	defer func() {
		_ = os.Remove("new.webp")
	}()

	f, err := os.Open(testPath)
	defer f.Close()

	m, err := Decode(f)

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("TestWebp_Encode Decode() returned %+v", err)
	}

	toimg, _ := os.Create("new.webp")
	defer toimg.Close()

	err = Encode(toimg, m, &Options{false, 50})

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("TestWebp_Encode Encode() returned %+v", err)
	}
}
