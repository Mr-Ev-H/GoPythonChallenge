package challenges

import (
	"fmt"
	"io/ioutil"
)

// Count character frequencies
// http://www.pythonchallenge.com/pc/def/ocr.html
func Solve2() {
	data, err := ioutil.ReadFile("./challenges/data/chal2input.txt")

	if err != nil {
		panic(err)
	}

	text := string(data)

	runeCount := make(map[rune]int)
	var seenRunes []rune

	for _, r := range text {
		seenRunes = append(seenRunes, r)
		runeCount[r]++
	}

	answer := ""

	for _, r := range seenRunes {
		if runeCount[r] == 1 {
			answer += string(r)
		}
	}

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/def/%s.html", answer)

	fmt.Println("Next: " + url)
}
