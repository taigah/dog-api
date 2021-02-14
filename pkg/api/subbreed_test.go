package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBreedsWithSubBreeds(t *testing.T) {
	breeds, err := GetAllBreedsWithSubBreeds()
	assert.Nil(t, err)
	subBreeds, ok := breeds["bullterrier"]
	assert.True(t, ok)
	assert.Contains(t, subBreeds, "staffordshire")
}

