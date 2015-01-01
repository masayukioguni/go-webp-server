package webp

import (
	"image"
	"io"

	ImageExt "github.com/chai2010/image"
	_ "github.com/chai2010/image/jpeg"
	_ "github.com/chai2010/image/webp"
)

type Options struct {
	Lossless bool
	Quality  float32
}

func Decode(r io.Reader) (image.Image, error) {
	m, _, err := ImageExt.Decode(r)

	if err != nil {
		return nil, err
	}

	return m, err

}

func Encode(w io.Writer, i image.Image, options *Options) error {
	return ImageExt.Encode("webp", w, i, ImageExt.NewOptions(options.Lossless, options.Quality))
}
