package challenges

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/h2non/filetype"
)

// First one I had real trouble with answering. The clue is a deck of cards with a barely readable 5 on it.
// The gfx file is the result of merging 5 images together, you have to unmerge them a byte at a time into
// 5 separate images.
// I added the file type checker so I could output the extension. Probably overkill, but I liked the interface
// to the library.
// http://www.pythonchallenge.com/pc/return/evil.html
func Solve12() {

	request, err := http.NewRequest(http.MethodGet, "http://www.pythonchallenge.com/pc/return/evil2.gfx", nil)

	if err != nil {
		panic(err)
	}

	request.SetBasicAuth("huge", "file")

	chalInput, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer chalInput.Body.Close()

	outImages := splitFiveWays(chalInput.Body)

	for i, image := range outImages {

		ft, err := filetype.Match(image)
		if err != nil {
			panic(err)
		}

		fileName := fmt.Sprintf("answers/challenge12_%d.%s", i+1, ft.Extension)
		file, err := os.Create(fileName)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		file.Write(image)

		fmt.Printf("Answer saved to: %s\n", fileName)
	}

}

// Given a reader, splits the data 1 byte at a time into 5 separate arrays
func splitFiveWays(reader io.Reader) [5][]byte {

	data, _ := ioutil.ReadAll(reader)

	output := [5][]byte{}

	for offset := 0; offset < 5; offset++ {

		output[offset] = make([]byte, 0)

		for i := offset; i < len(data); i += 5 {
			output[offset] = append(output[offset], data[i])
		}
	}

	return output
}
