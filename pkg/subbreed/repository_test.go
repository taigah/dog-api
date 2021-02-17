package subbreed

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_SubBreeds_Of_Should_Return_The_Sub_Breeds_Of_A_Given_Breed(t *testing.T) {
	rep := NewRepository(http.DefaultClient)
	subBreeds, err := rep.GetSubBreedsOf("bullterrier")
	assert.Nil(t, err)
	assert.Contains(t, subBreeds, "staffordshire")
}
