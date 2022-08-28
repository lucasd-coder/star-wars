package http_client

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseFromHttpResponse(resp *http.Response, model interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.Unmarshal(bodyBytes, &model)

	if err != nil {
		return err
	}

	return nil
}
