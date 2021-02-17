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

func Test_GetBunchRandoms_Should_Return_A_Bunch_Of_Random_Images(t *testing.T) {
	rep := NewImageRepository(http.DefaultClient)
	imageCount := 30
	images, err := rep.GetBunchRandoms(imageCount)
	assert.Nil(t, err)
	assert.Len(t, images, imageCount)
	for _, image := range images {
		assert.True(t, strings.HasPrefix(image, "http"))
	}
}

func Test_GetAllByBreed_Should_Return_All_Images_Of_The_Given_Breed(t *testing.T) {
	rep := NewImageRepository(http.DefaultClient)
	breed := "hound"
	images, err := rep.GetAllByBreed(breed)
	assert.Nil(t, err)
	assert.Contains(t, images, "https://images.dog.ceo/breeds/hound-afghan/n02088094_1003.jpg")
	for _, image := range images {
		assert.True(t, strings.HasPrefix(image, "http"))
	}
}

func Test_GetRandomByBreed_Should_Return_A_Random_Image_Of_The_Given_Breed(t *testing.T) {
	rep := NewImageRepository(http.DefaultClient)
	breed := "hound"
	image, err := rep.GetRandomByBreed(breed)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(image, "http"))
}
