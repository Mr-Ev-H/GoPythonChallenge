package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

// Get the size of the 30'th item in the look and say sequence.
// Enjoted the algorithm bit of it, although I had to look
// up the sequence from https://mathworld.wolfram.com/LookandSaySequence.html
// http://www.pythonchallenge.com/pc/return/bull.html
func Solve10() {

	a := make([]string, 31)

	a[0] = "1"
	lasVal := a[0]

	for i := 1; i < 31; i++ {
		lasVal = lookAndSay(lasVal)
		a[i] = lasVal
	}

	answer := len(a[30])

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/return/%d.html", answer)
	fmt.Println("Next: " + url)

}

func lookAndSay(number string) string {

	b := strings.Builder{}

	for i := 0; i < len(number); i++ {
		char := string(number[i])
		currentVal, _ := strconv.Atoi(char)
		nextVal := currentVal
		count := int64(1)

		nexti := i + 1

		for nexti < len(number) && nextVal == currentVal {
			nextVal, _ = strconv.Atoi(string(number[nexti]))
			if nextVal == currentVal {
				count++
				i++
			}

			nexti++
		}

		b.WriteString(strconv.Itoa(int(count)))
		b.WriteString(strconv.Itoa(int(currentVal)))
	}

	return b.String()
}
