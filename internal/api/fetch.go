package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
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

func buildURL(endpoint string) string {
	u, err := url.Parse("https://dog.ceo/api/")
	if err != nil {
		log.Fatalf("Could not build '%v'", endpoint)
	}
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

func Fetch(client *http.Client, endpoint string, data interface{}) error {
	url := buildURL(endpoint)
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
