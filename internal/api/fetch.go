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

func Fetch(client *http.Client, endpoint string, data interface{}) error {
	url := fmt.Sprintf("https://dog.ceo/api/%v", endpoint)
	res, err := client.Get(url)

	if err != nil {
		return err
	}

	if res.StatusCode > 300 {
		var err struct {
			Message string `json:"message"`
		}

		unmarshalResponse(res, &err)

		return errors.New(err.Message)
	}

	unmarshalResponse(res, &data)
	return nil
}
