package utils

import (
	"fmt"
	"log"
	"net/http"

	"io/ioutil"
)

func RequestGet(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Error: %s", body)
		return nil, fmt.Errorf("Error: %s", body)
	}

	return []byte{}, nil

}
