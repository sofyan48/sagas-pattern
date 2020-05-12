package requester

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type Requester struct{}

func RequesterHandler() *Requester {
	return &Requester{}
}

type RequesterInterface interface {
	GET(url, authorization string) ([]byte, error)
	POST(url, auth string, payload []byte) ([]byte, error)
}

// GET request type get
func (request *Requester) GET(url, authorization string) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}

	return body, nil
}

// POST request type post
func (request *Requester) POST(url, auth string, payload []byte) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Keep-Alive", "timeout=1, max=1000")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	return body, nil
}
