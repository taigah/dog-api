package image

import (
	"fmt"
	"net/http"

	ihttp "github.com/taigah/dog-api/internal/http"
	"github.com/taigah/dog-api/pkg/breed"
	"github.com/taigah/dog-api/pkg/subbreed"
)

type Repository interface {
	GetRandom() (Image, error)
	GetBunchRandoms(imageCount int) ([]Image, error)

	GetAllByBreed(breed breed.Breed) ([]Image, error)
	GetRandomByBreed(breed breed.Breed) (Image, error)
	GetBunchRandomsByBreed(breed breed.Breed, imageCount int) ([]Image, error)

	GetAllBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed) ([]Image, error)
	GetRandomBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed) (Image, error)
	GetBunchRandomsBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed, imageCount int) ([]Image, error)
}

type imageRepositoryImpl struct {
	client *http.Client
}

func NewRepository(client *http.Client) Repository {
	return &imageRepositoryImpl{client}
}

func (rep imageRepositoryImpl) getImageFromURL(url string) (Image, error) {
	var data struct {
		Image string `json:"message"`
	}

	err := ihttp.UnmarshalFromURL(rep.client, url, &data)

	if err != nil {
		return "", err
	}

	return data.Image, err
}

func (rep imageRepositoryImpl) getImagesFromURL(url string) ([]Image, error) {
	var data struct {
		Images []string `json:"message"`
	}

	err := ihttp.UnmarshalFromURL(rep.client, url, &data)

	if err != nil {
		return nil, err
	}

	return data.Images, err
}

func (rep imageRepositoryImpl) GetRandom() (Image, error) {
	return rep.getImageFromURL("https://dog.ceo/api/breeds/image/random")
}

func (rep imageRepositoryImpl) GetBunchRandoms(imageCount int) ([]Image, error) {
	return rep.getImagesFromURL(fmt.Sprintf("https://dog.ceo/api/breeds/image/random/%v", imageCount))
}

func (rep imageRepositoryImpl) GetAllByBreed(breed breed.Breed) ([]Image, error) {
	return rep.getImagesFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/images", breed))
}

func (rep imageRepositoryImpl) GetRandomByBreed(breed breed.Breed) (Image, error) {
	return rep.getImageFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/images/random", breed))
}

func (rep imageRepositoryImpl) GetBunchRandomsByBreed(breed breed.Breed, imageCount int) ([]Image, error) {
	return rep.getImagesFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/images/random/%v", breed, imageCount))
}

func (rep imageRepositoryImpl) GetAllBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed) ([]Image, error) {
	return rep.getImagesFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/%v/images", breed, subBreed))
}

func (rep imageRepositoryImpl) GetRandomBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed) (Image, error) {
	return rep.getImageFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/%v/images/random", breed, subBreed))
}

func (rep imageRepositoryImpl) GetBunchRandomsBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed, imageCount int) ([]Image, error) {
	return rep.getImagesFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/%v/images/random/%v", breed, subBreed, imageCount))
}
