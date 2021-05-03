package challenges

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

// Given two interleaved images in a single jpeg, seperate them out from one another and write result to file.
// http://www.pythonchallenge.com/pc/return/5808.html
func Solve11() {

	request, err := http.NewRequest(http.MethodGet, "http://www.pythonchallenge.com/pc/return/cave.jpg", nil)

	if err != nil {
		panic(err)
	}

	request.SetBasicAuth("huge", "file")

	chalImage, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer chalImage.Body.Close()

	imageData, err := jpeg.Decode(chalImage.Body)

	if err != nil {
		panic(err)
	}

	bounds := image.Rect(0, 0, 640, 480)

	evenImage := image.NewRGBA(bounds)

	for x := 0; x < imageData.Bounds().Dx(); x++ {
		for y := 0; y < imageData.Bounds().Dy(); y++ {
			if x%2 == 0 && y%2 == 0 {
				evenImage.Set(x, y, imageData.At(x, y))
			} else {
				evenImage.Set(x, y, color.Black)
			}
		}
	}

	outFileEven, err := os.Create("answers/chal11.png")

	if err != nil {
		panic(err)
	}

	defer outFileEven.Close()

	png.Encode(outFileEven, evenImage)

	fmt.Println("Answer saved to: " + outFileEven.Name())
}
