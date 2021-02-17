package image

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Random_Should_Return_A_Random_Image(t *testing.T) {
	rep := NewImageRepository(http.DefaultClient)
	image, err := rep.GetRandom()
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(image, "http"))
}
