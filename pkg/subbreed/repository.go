package subbreed

import (
	"fmt"
	"net/http"

	httpi "github.com/taigah/dog-api/internal/http"
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

	err := httpi.UnmarshalFromURL(rep.client, fmt.Sprintf("https://dog.ceo/api/breed/%v/list", breed), &data)

	if err != nil {
		return nil, err
	}

	return data.SubBreeds, nil
}
