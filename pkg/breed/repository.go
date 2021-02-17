package breed

import (
	"net/http"

	"github.com/taigah/dog-api/internal/api"
)

type Repository interface {
	GetAll() ([]Breed, error)
}

type repo struct {
	client *http.Client
}

func NewRepository(client *http.Client) Repository {
	return &repo{client}
}

func (rep *repo) GetAll() ([]Breed, error) {
	var data struct {
		Breeds map[string][]string `json:"message"`
	}

	err := api.Fetch(rep.client, "breeds/list/all", &data)

	if err != nil {
		return nil, err
	}

	breeds := make([]Breed, 0)

	for breed, _ := range data.Breeds {
		breeds = append(breeds, breed)
	}

	return breeds, nil
}
