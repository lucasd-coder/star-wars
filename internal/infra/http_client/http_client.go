package http_client

import (
	"net/http"
)

func Get(url string, headers map[string]string, t *http.Transport) (*http.Response, error) {
	httpClient := &http.Client{
		Transport: t,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
