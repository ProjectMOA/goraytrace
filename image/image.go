package image

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ProjectMOA/goraytrace/math3d"
)

// Image defines an RGB image
type Image struct {
	Width, Height int32
	Pixels        [][]Color
}

// New returns a new Image with the pixel slice initialized
func New(width int32, height int32) *Image {
	retVal := &Image{Width: width, Height: height}
	retVal.Pixels = make([][]Color, height, height)
	for i := int32(0); i < height; i++ {
		retVal.Pixels[i] = make([]Color, width, width)
	}
	return retVal
}

// Save saves the image as a ppm image.
func (img *Image) Save(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("P3\n")
	writer.WriteString("# " + filename + "\n")
	writer.WriteString(fmt.Sprint(img.Width) + " " + fmt.Sprint(img.Height) + "\n")
	writer.WriteString("255\n")

	// Shorthand for clamping bytes
	bClamp := func(v float64) uint8 {
		return uint8(math3d.Clamp(v, 0, 255))
	}

	for i := 0; i < len(img.Pixels); i++ {
		for j := 0; j < len(img.Pixels[0]); j++ {
			var tc = img.Pixels[i][j]
			writer.WriteString(
				fmt.Sprintf("%d %d %d ",
					bClamp(tc.R*255),
					bClamp(tc.G*255),
					bClamp(tc.B*255)))
		}
		writer.WriteString("\n")
	}

}
