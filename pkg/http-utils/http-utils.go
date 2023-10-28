package httputils

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetData(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}

func GetJSON[T interface{}](url string) (*T, error) {
	data, err := GetData(url)

	if err != nil {
		return nil, err
	}

	var jsonData T

	err = json.Unmarshal(data, &jsonData)

	if err != nil {
		return nil, err
	}

	return &jsonData, nil
}
