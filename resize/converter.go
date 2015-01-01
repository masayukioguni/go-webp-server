package resize

import (
	libResize "github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
)

func Resize(r io.Reader, width uint, height uint) (image.Image, error) {
	img, err := jpeg.Decode(r)
	if err != nil {
		return nil, err
	}

	m := libResize.Resize(width, height, img, libResize.Lanczos3)

	return m, nil

}
