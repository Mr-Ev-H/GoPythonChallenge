package challenges

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"gopychallenge/utility"
	"io"
	"net/http"
)

// Linked list with a zip + reading file info headers
// closure, zip + function abstraction
// Type assertions, debugging, module imports  - weird
// http://www.pythonchallenge.com/pc/def/channel.html
func Solve6() {

	zipRequest, err := http.Get("http://www.pythonchallenge.com/pc/def/channel.zip")

	if err != nil {
		panic(err)
	}

	zipBytes, err := io.ReadAll(zipRequest.Body)

	if err != nil {
		panic(err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), zipRequest.ContentLength)

	if err != nil {
		panic(err)
	}

	comments := make([]string, 10)

	navFunction := func(navigateId string) string {
		nextId, comment := zipNavigate(zipReader, navigateId)

		comments = append(comments, comment)
		return nextId
	}

	utility.ListWalker(navFunction, 90052)

	for _, c := range comments {
		fmt.Print(c)
	}

	fmt.Println("Next: http://www.pythonchallenge.com/pc/def/☝️ .html")
}

func zipNavigate(zipReader *zip.Reader, navigateId string) (string, string) {
	fmt.Println(navigateId)

	file, err := zipReader.Open(navigateId + ".txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := ""

	for scanner.Scan() {
		output += scanner.Text()
	}

	stat, statErr := file.Stat()

	if statErr != nil {
		panic(statErr)
	}

	// Zip.FileInfoHeader doesnt populate comments, casting to a file header gets us access to them
	fsys, ok := stat.Sys().(*zip.FileHeader)

	if !ok {
		panic(errors.New("unable to load compressed file header"))
	}

	return output, fsys.Comment
}
