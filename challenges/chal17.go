package challenges

import (
	"fmt"
	"gopychallenge/utility"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Closure over funcs
// Basic cookie reading
// Rehashing of linked list logic + calling phonebook rpc
// + sending cookie values
// This is one of the more annoying challenges
// http://www.pythonchallenge.com/pc/def/romance.html
func Solve17() {

	cookieData := ""

	navWrapper := func(navigateId string) string {
		next, cookieValue := extractNextAndCookieData(navigateId)
		value, _ := url.QueryUnescape(cookieValue)
		cookieData += value
		return next
	}

	utility.ListWalker(navWrapper, 12345)

	clue := readbzip(cookieData)

	fmt.Println("Clue 1: " + clue)

	fmt.Println("Making Phonebook RPC Call for Leopold (Mozarts Father)...")

	answer, _ := LookupNumber("Leopold")

	answer = strings.Trim(answer, "5-")
	answer = strings.ToLower(answer)

	fmt.Println("Clue 2: " + answer)

	//Then you have do some other fiddling around to figure out that you have to submit the message to the violin page with
	// The message in the cookie header

	answerUrl := fmt.Sprintf("http://www.pythonchallenge.com/pc/stuff/%s.php", answer)

	request, _ := http.NewRequest(http.MethodPost, answerUrl, nil)

	cookieValue := url.QueryEscape("the flowers are on their way")
	cookie := http.Cookie{Name: "info", Value: cookieValue}
	request.AddCookie(&cookie)

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	bodyString := string(body)

	fmt.Println(bodyString)
}

func extractNextAndCookieData(navigateId string) (string, string) {

	request, _ := http.NewRequest(http.MethodGet, "http://www.pythonchallenge.com/pc/def/linkedlist.php", nil)
	request.URL.ForceQuery = true

	q := request.URL.Query()
	q.Add("busynothing", navigateId)
	request.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		panic(readErr)
	}

	bodyString := string(body)

	cookies := resp.Cookies()

	return bodyString, cookies[0].Value
}
