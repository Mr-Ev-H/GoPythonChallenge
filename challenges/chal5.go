package challenges

import (
	"fmt"
	"gopychallenge/utility"
	"strings"

	"github.com/hydrogen18/stalecucumber"
)

// Python pickler
// Thanks to hydrogen18 didnt have to write the depickler myself which would have been miserable
// Type assertions, debugging, module imports
// http://www.pythonchallenge.com/pc/def/peak.html
func Solve5() {

	pickedData, err := utility.GetFileForChallenge("http://www.pythonchallenge.com/pc/def/banner.p", "", "")

	if err != nil {
		panic(err)
	}

	var output [][][]interface{} //The shape of a python tuple

	pickleErr := stalecucumber.UnpackInto(&output).From(stalecucumber.Unpickle(pickedData))

	if pickleErr != nil {
		panic(pickleErr)
	}

	fmt.Println("Hint:")
	for _, line := range output {
		for _, segments := range line {
			character := segments[0].(string)
			count := segments[1].(int64)
			fmt.Print(strings.Repeat(character, int(count)))
		}

		fmt.Println()
	}

	fmt.Println("Next: http://www.pythonchallenge.com/pc/def/☝️ .html")
}
