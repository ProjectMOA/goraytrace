package image

import (
	stdimg "image"
	"image/png"
	"log"
	"os"
)

// Image defines an RGB image
type Image struct {
	stdimg.NRGBA
}

// New returns a new Image with the pixel slice initialized
func New(width int, height int) *Image {
	return &Image{*stdimg.NewNRGBA(stdimg.Rect(0, 0, width, height))}
}

// Save saves the image as a ppm image.
func (img *Image) Save(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	f, err := os.Create(filename + ".png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
