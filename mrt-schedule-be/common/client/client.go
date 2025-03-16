package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoRequest(client *http.Client, url string)([]byte, error){
	resp, err := client.Get(url)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected Status Code: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	return body, nil
}