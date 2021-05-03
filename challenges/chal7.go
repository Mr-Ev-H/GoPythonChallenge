package challenges

import (
	"fmt"
	"image/png"
	"net/http"
	"strconv"
	"strings"
)

// Given an image with a stripe of greyscale pixels, read value of pixels, convert to bytes, convert bytes to strings,
// then convert internal array to more bytes to convert into the answer string.
// http://www.pythonchallenge.com/pc/def/oxygen.html
func Solve7() {

	chalImage, err := http.Get("http://www.pythonchallenge.com/pc/def/oxygen.png")

	if err != nil {
		panic(err)
	}

	defer chalImage.Body.Close()

	imageData, decodeErr := png.Decode(chalImage.Body)

	if decodeErr != nil {
		panic(decodeErr)
	}

	s := make([]byte, 0)

	lineWidth := 607
	pixelSize := 7

	for i := 0; i < lineWidth; i += pixelSize {
		// fmt.Print(imageData.At(i, 50))
		d := imageData.At(i, 50)
		r, _, _, _ := d.RGBA()
		s = append(s, byte(r))
	}

	output := string(s)

	fmt.Println(output)

	arrIndexStart := strings.Index(output, "[")
	arrIndexEnd := strings.Index(output, "]")

	array := output[arrIndexStart+1 : arrIndexEnd]

	csv := strings.Trim(array, " ")

	arr := strings.Split(csv, ",")

	answer := ""

	for _, c := range arr {
		val, _ := strconv.Atoi(strings.Trim(c, " "))

		answer += string(byte(val))
	}

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/def/%s.html", answer)
	fmt.Println("Next: " + url)
}
