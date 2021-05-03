package challenges

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Another fiddly challenge. Working with the range header in http
// Need to get a set of
// http://www.pythonchallenge.com/pc/hex/idiot2.html
func Solve20() {
	buf := new(strings.Builder)

	nextRange := 30203 // First request for the image tells us that we need to start from this offset

	for nextRange > 0 {
		var body io.Reader
		nextRange, body = readBytesFromRange(nextRange)

		if nextRange > 0 {
			io.Copy(buf, body)
			fmt.Println(buf.String())
			buf.Reset()
		}
	}

	//Request for the end of the stream
	_, body := readBytesFromRange(2123456789)
	io.Copy(buf, body)
	fmt.Println(buf.String())
	buf.Reset()

	//Request for 1 byte before range returned
	_, body = readBytesFromRange(2123456743)
	io.Copy(buf, body)
	fmt.Println(buf.String())

	_, responseBody := readBytesFromRange(1152983631)

	outputFile, _ := os.Create("answers/chal20.zip")

	io.Copy(outputFile, responseBody)

	fmt.Println("Answer saved to: " + outputFile.Name())
	fmt.Println("Password: redavni")
}

//Reads the range offset header and returns the start of the next byte sequence to request
func readBytesFromRange(startBytes int) (int, io.Reader) {
	request, _ := http.NewRequest(http.MethodGet, "http://www.pythonchallenge.com/pc/hex/unreal.jpg", nil)
	request.SetBasicAuth("butter", "fly")

	request.Header.Add("Range", fmt.Sprintf("bytes=%d-", startBytes))

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println(err)
	} else {
		rangeReg := regexp.MustCompile(`(\d+)-(\d+)/(\d+)`)
		rangeHeader := resp.Header.Get("Content-Range")
		rangeHeader = strings.Trim(rangeHeader, "bytes ")

		rangeBytes := rangeReg.FindStringSubmatch(rangeHeader)

		if len(rangeBytes) > 0 {
			fmt.Printf("Range From: %s, To: %s\n", rangeBytes[1], rangeBytes[2])

			val, _ := strconv.Atoi(rangeBytes[2])
			return val + 1, resp.Body
		}

	}

	return 0, nil
}
