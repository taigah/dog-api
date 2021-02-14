package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBreeds(t *testing.T) {
	breeds, err := GetAllBreeds()
	assert.Nil(t, err)
	assert.Contains(t, breeds, "dingo")
}
