package challenges

import (
	"fmt"
	"math"
)

// Basic math
// http://www.pythonchallenge.com/pc/def/0.html
func Solve0() {
	answer := (int)(math.Pow(2, 38))
	fmt.Printf("2^38=%d\n", answer)

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/def/%d.html", answer)

	fmt.Println("Next: " + url)

}
