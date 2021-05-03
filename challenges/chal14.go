package challenges

import (
	"fmt"
	"gopychallenge/utility"
	"image"
	"image/png"
	"os"
)

// Given a 1d array, unwind clockwise into an image and output
// More fiddly than difficult.
// http://www.pythonchallenge.com/pc/return/italy.html
func Solve14() {

	chalImage, err := utility.GetFileForChallenge("http://www.pythonchallenge.com/pc/return/wire.png", "huge", "file")

	if err != nil {
		panic(err)
	}

	img, err := png.Decode(chalImage)

	if err != nil {
		panic(err)
	}

	bounds := image.Rect(0, 0, 100, 100)

	outImage := image.NewRGBA(bounds)

	remainingPixels := bounds.Max.X * bounds.Max.Y
	pixel := 0

	// Each time we write a full line, the boundaries get closer
	fromX := 0
	fromY := 0
	toX := 99
	toY := 99

	for pixel < remainingPixels {

		for x := fromX; x <= toX; x++ {
			color := img.At(pixel, 0)
			outImage.Set(x, fromY, color)
			pixel++
		}

		fromY++

		for y := fromY; y <= toY; y++ {
			color := img.At(pixel, 0)
			outImage.Set(toX, y, color)
			pixel++
		}

		toX--

		for x := toX; x >= fromX; x-- {
			color := img.At(pixel, 0)
			outImage.Set(x, toY, color)
			pixel++
		}

		toY--

		for y := toY; y >= fromY; y-- {
			color := img.At(pixel, 0)
			outImage.Set(fromX, y, color)
			pixel++
		}

		fromX++
	}

	outFile, err := os.Create("answers/chal14.png")

	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	png.Encode(outFile, outImage)

	fmt.Println("Answer saved to: " + outFile.Name())
}
