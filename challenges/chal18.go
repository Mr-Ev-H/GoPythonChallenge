package challenges

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"gopychallenge/utility"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/pmezard/go-difflib/difflib"
)

// Diffing text files
// Quite a bit of a pain, tried to handle it via images, then naievely, then via a different diffing library, managed to get it working with the partial
// Implementation of python difflib but its a massive kludge and I hate it. I'd like to know whether theres a better way of solving this.
// http://www.pythonchallenge.com/pc/return/balloons.html
func Solve18() {

	chalInputZip, err := utility.GetFileForChallenge("http://www.pythonchallenge.com/pc/return/deltas.gz", "huge", "file")

	if err != nil {
		panic(err)
	}

	zipBytes, err := io.ReadAll(chalInputZip)

	if err != nil {
		panic(err)
	}

	zipReader, err := gzip.NewReader(bytes.NewReader(zipBytes))

	if err != nil {
		panic(err)
	}

	// file, err := zipReader.("delta.txt")

	// if err != nil {
	// 	panic(err)
	// }

	scanner := bufio.NewScanner(zipReader)

	leftData := make([]string, 0)
	rightData := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		left := line[:53] + "\n" // Although diff lib takes line arrays, it only works line by line if theyre separated by newline
		right := line[56:] + "\n"

		leftData = append(leftData, left)
		rightData = append(rightData, right)
	}

	bothFile, _ := os.Create("answers/chal18_both.png")
	leftFile, _ := os.Create("answers/chal18_left.png")
	rightFile, _ := os.Create("answers/chal18_right.png")

	defer bothFile.Close()
	defer leftFile.Close()
	defer rightFile.Close()

	diff := difflib.UnifiedDiff{
		A:        leftData,
		B:        rightData,
		FromFile: "Left",
		ToFile:   "Right",
		Context:  10000, //This sucks - pythons difflib actually does have support for outputting total diffs, but the go implementation does not
	}

	result, _ := difflib.GetUnifiedDiffString(diff)

	lines := difflib.SplitLines(result)

	for _, line := range lines[3:] {
		line := strings.TrimSuffix(line, "\n")
		diffPrefix, _ := utf8.DecodeRuneInString(line)

		if diffPrefix == ' ' {
			bothFile.Write(convertStringToBinary(line[1:]))
		} else if diffPrefix == '+' {
			rightFile.Write(convertStringToBinary(line[1:]))
		} else if diffPrefix == '-' {
			leftFile.Write(convertStringToBinary(line[1:]))
		}
	}

	fmt.Println("Answer saved to: " + bothFile.Name())
	fmt.Println("Answer saved to: " + leftFile.Name())
	fmt.Println("Answer saved to: " + rightFile.Name())
}

func convertStringToBinary(byteString string) []byte {
	bsArray := strings.Split(byteString, " ")

	lineBytes := make([]byte, 0)

	for _, bs := range bsArray {
		b, _ := strconv.ParseInt(bs, 16, 0)
		lineBytes = append(lineBytes, byte(b))
	}

	return lineBytes
}
