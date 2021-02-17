package subbreed

import (
	"fmt"
	"net/http"

	"github.com/taigah/dog-api/internal/api"
	"github.com/taigah/dog-api/pkg/breed"
)

type Repository interface {
	GetSubBreedsOf(breed breed.Breed) ([]SubBreed, error)
}

type repositoryImpl struct {
	client *http.Client
}

func NewRepository(client *http.Client) Repository {
	return &repositoryImpl{client}
}

func (rep *repositoryImpl) GetSubBreedsOf(breed breed.Breed) ([]SubBreed, error) {
	var data struct {
		SubBreeds []string `json:"message"`
	}

	err := api.Fetch(rep.client, fmt.Sprintf("breed/%v/list", breed), &data)

	if err != nil {
		return nil, err
	}

	return data.SubBreeds, nil
}
