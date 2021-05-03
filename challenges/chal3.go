package challenges

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// Pattern matching using regex
// http://www.pythonchallenge.com/pc/def/equality.html
func Solve3() {
	data, err := ioutil.ReadFile("./challenges/data/chal3input.txt")

	if err != nil {
		panic(err)
	}

	text := string(data)

	reg := regexp.MustCompile(`(?:[a-z^][A-Z]{3})([a-z])(?:[A-Z]{3}[a-z$])`)

	matches := reg.FindAllStringSubmatch(text, -1)

	answer := ""

	for i := range matches {
		answer += matches[i][1]
	}

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/def/%s.html", answer)

	fmt.Println("Next: " + url)
}

// Pattern matching using dynamic programming
// I didn't spot the re hint in the title and chose to solve via dyanmic programming instead of regexes. See Solve3 for regex solution
// For obvious reasons this is signitifantly more performant than the regex solution
// http://www.pythonchallenge.com/pc/def/equality.html
func Solve3Dyn() {
	data, err := ioutil.ReadFile("./challenges/data/chal3input.txt")

	if err != nil {
		panic(err)
	}

	text := string(data)

	var output []rune

	upCount := 0
	downCount := 0

	var candidate rune

	for _, r := range text {
		if r >= 'A' && r <= 'Z' {
			upCount++
			if candidate != 0 {
				downCount++
			}
		} else if r >= 'a' && r <= 'z' {

			if upCount == 3 && downCount == 3 {
				output = append(output, candidate)
			}

			if candidate != 0 {
				candidate = 0
			}

			if upCount == 3 {
				candidate = r
			}

			upCount = 0
			downCount = 0
		}
	}

	answer := string(output)

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/def/%s.html", answer)

	fmt.Println("Next: " + url)
}
