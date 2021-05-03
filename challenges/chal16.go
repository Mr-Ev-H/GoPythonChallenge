package challenges

import (
	"fmt"
	"gopychallenge/utility"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
)

// Shuffle lines in an image
// More Array and pixel manipulation + a divmod
// http://www.pythonchallenge.com/pc/return/mozart.html
func Solve16() {
	chalImage, err := utility.GetFileForChallenge("http://www.pythonchallenge.com/pc/return/mozart.gif", "huge", "file")

	if err != nil {
		panic(err)
	}

	img, err := gif.Decode(chalImage)

	if err != nil {
		panic(err)
	}

	outImage := image.NewRGBA(img.Bounds())
	mysteryPink := color.RGBA{R: 255, G: 0, B: 255, A: 255}

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {

			if img.At(x, y) == mysteryPink {

				for i := 0; i < img.Bounds().Max.X; i++ {
					fromX := (i + x) % img.Bounds().Max.X
					fromImage := img.At(fromX, y)
					outImage.Set(i, y, fromImage)
				}
			}
		}
	}

	outFile, err := os.Create("answers/chal16.png")

	if err != nil {
		panic(err)
	}

	png.Encode(outFile, outImage)

	fmt.Println("Answer saved to: " + outFile.Name())
}
