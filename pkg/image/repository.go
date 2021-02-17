package image

import (
	"net/http"

	ihttp "github.com/taigah/dog-api/internal/http"
)

type ImageRepository interface {
	GetRandom() (Image, error)
}

type imageRepositoryImpl struct {
	client *http.Client
}

func NewImageRepository(client *http.Client) ImageRepository {
	return &imageRepositoryImpl{client}
}

func (rep imageRepositoryImpl) GetRandom() (Image, error) {
	var data struct {
		Image string `json:"message"`
	}
	err := ihttp.UnmarshalFromURL(rep.client, "https://dog.ceo/api/breeds/image/random", &data)

	if err != nil {
		return "", err
	}

	return data.Image, nil
}
