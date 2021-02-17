package image

import (
	"fmt"
	"net/http"

	"github.com/taigah/dog-api/internal/api"
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

func (rep imageRepositoryImpl) getImage(endpoint string) (Image, error) {
	var data struct {
		Image string `json:"message"`
	}

	err := api.Fetch(rep.client, endpoint, &data)

	if err != nil {
		return "", err
	}

	return data.Image, err
}

func (rep imageRepositoryImpl) getImages(endpoint string) ([]Image, error) {
	var data struct {
		Images []string `json:"message"`
	}

	err := api.Fetch(rep.client, endpoint, &data)

	if err != nil {
		return nil, err
	}

	return data.Images, err
}

func (rep imageRepositoryImpl) GetRandom() (Image, error) {
	return rep.getImage("breeds/image/random")
}

func (rep imageRepositoryImpl) GetBunchRandoms(imageCount int) ([]Image, error) {
	return rep.getImages(fmt.Sprintf("breeds/image/random/%v", imageCount))
}

func (rep imageRepositoryImpl) GetAllByBreed(breed breed.Breed) ([]Image, error) {
	return rep.getImages(fmt.Sprintf("breed/%v/images", breed))
}

func (rep imageRepositoryImpl) GetRandomByBreed(breed breed.Breed) (Image, error) {
	return rep.getImage(fmt.Sprintf("breed/%v/images/random", breed))
}

func (rep imageRepositoryImpl) GetBunchRandomsByBreed(breed breed.Breed, imageCount int) ([]Image, error) {
	return rep.getImages(fmt.Sprintf("breed/%v/images/random/%v", breed, imageCount))
}

func (rep imageRepositoryImpl) GetAllBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed) ([]Image, error) {
	return rep.getImages(fmt.Sprintf("breed/%v/%v/images", breed, subBreed))
}

func (rep imageRepositoryImpl) GetRandomBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed) (Image, error) {
	return rep.getImage(fmt.Sprintf("breed/%v/%v/images/random", breed, subBreed))
}

func (rep imageRepositoryImpl) GetBunchRandomsBySubBreed(breed breed.Breed, subBreed subbreed.SubBreed, imageCount int) ([]Image, error) {
	return rep.getImages(fmt.Sprintf("breed/%v/%v/images/random/%v", breed, subBreed, imageCount))
}
