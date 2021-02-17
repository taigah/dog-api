package image

import (
	"fmt"
	"net/http"

	ihttp "github.com/taigah/dog-api/internal/http"
)

type ImageRepository interface {
	GetRandom() (Image, error)
	GetBunchRandoms(imageCount int) ([]Image, error)
}

type imageRepositoryImpl struct {
	client *http.Client
}

func NewImageRepository(client *http.Client) ImageRepository {
	return &imageRepositoryImpl{client}
}

func (rep imageRepositoryImpl) GetRandom() (Image, error) {
	var data struct {
		Image Image `json:"message"`
	}
	err := ihttp.UnmarshalFromURL(rep.client, "https://dog.ceo/api/breeds/image/random", &data)

	if err != nil {
		return "", err
	}

	return data.Image, nil
}

func (rep imageRepositoryImpl) GetBunchRandoms(imageCount int) ([]Image, error) {
	var data struct {
		Images []Image `json:"message"`
	}
	err := ihttp.UnmarshalFromURL(rep.client, fmt.Sprintf("https://dog.ceo/api/breeds/image/random/%v", imageCount), &data)

	if err != nil {
		return nil, err
	}

	return data.Images, nil
}
