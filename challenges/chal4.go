package challenges

import (
	"fmt"
	"gopychallenge/utility"
	"io"
	"net/http"
)

// Linked list of URLs that they've helpfully thrown some instructions into that you have to parse out.
// Separated the logic into the ListWalker function as we have to use this a few times throuhgout the challenge
// http://www.pythonchallenge.com/pc/def/linkedlist.html
func Solve4() {

	answer := utility.ListWalker(navigate, 0)

	url := fmt.Sprintf("http://www.pythonchallenge.com/pc/def/%s.html", answer)

	fmt.Println("Next: " + url)
}

func navigate(navigateId string) string {
	resp, err := http.Get("http://www.pythonchallenge.com/pc/def/linkedlist.php?nothing=" + navigateId)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		panic(readErr)
	}

	return string(body)
}
