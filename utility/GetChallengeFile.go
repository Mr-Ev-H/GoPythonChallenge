package utility

import (
	"io"
	"net/http"
)

func GetFileForChallenge(fileUrl string, username string, password string) (io.Reader, error) {
	request, err := http.NewRequest(http.MethodGet, fileUrl, nil)

	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(username, password)

	chalInput, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	return chalInput.Body, nil
}
