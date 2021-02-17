package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func unmarshalResponse(res *http.Response, data interface{}) error {
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(body, &data)
}

func isErrorResponse(res *http.Response) bool {
	return res.StatusCode >= 400
}

func Fetch(client *http.Client, endpoint string, data interface{}) error {
	url := fmt.Sprintf("https://dog.ceo/api/%v", endpoint)
	res, err := client.Get(url)

	if err != nil {
		return err
	}

	if isErrorResponse(res) {
		var apiError struct {
			Message string `json:"message"`
		}

		err := unmarshalResponse(res, &apiError)

		if err != nil {
			return err
		}

		return errors.New(apiError.Message)
	}

	return unmarshalResponse(res, data)
}
