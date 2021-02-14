package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func unmarshalFromURL(url string, data interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}
