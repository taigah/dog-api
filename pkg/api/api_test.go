package api

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBreeds(t *testing.T) {
	breeds, err := GetAllBreeds()
	assert.Nil(t, err)
	assert.Contains(t, breeds, "dingo")
}

func TestGetRandomDogImageUrl(t *testing.T) {
	url, err := GetRandomDogImageUrl()
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(url, "http"))
}

func TestGetBunchOfRandomDogImagesUrl(t *testing.T) {
	imageCount := 50
	urls, err := GetBunchOfRandomDogImagesUrl(imageCount)
	assert.Nil(t, err)
	for _, url := range urls {
		assert.True(t, strings.HasPrefix(url, "http"))
	}
	assert.Equal(t, len(urls), imageCount)
}

func TestGetAllDogsOfBreed(t *testing.T) {
	urls, err := GetAllDogsOfBreed("hound")
	assert.Nil(t, err)
	for _, url := range urls {
		assert.True(t, strings.HasPrefix(url, "http"))
	}
	assert.Contains(t, urls, "https://images.dog.ceo/breeds/hound-afghan/n02088094_1003.jpg")
}

func TestGetRandomDogOfBreed(t *testing.T) {
	url, err := GetRandomDogOfBreed("hound")
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(url, "http"))
}

func TestGetBunchOfRandomDogsOfBreed(t *testing.T) {
	imageCount := 50
	urls, err := GetBunchOfRandomDogsOfBreed("hound", imageCount)
	assert.Nil(t, err)
	for _, url := range urls {
		assert.True(t, strings.HasPrefix(url, "http"))
	}
	assert.Equal(t, len(urls), imageCount)
}

func TestGetSubBreeds(t *testing.T) {
	subBreeds, err := GetSubBreeds("hound")
	assert.Nil(t, err)
	assert.Contains(t, subBreeds, "afghan")
}

func TestGetDogOfSubBreed(t *testing.T) {
	urls, err := GetDogsOfSubBreed("hound", "afghan")
	assert.Nil(t, err)
	assert.Contains(t, urls, "https://images.dog.ceo/breeds/hound-afghan/n02088094_1186.jpg")
}

func TestGetRandomDogOfSubBreed(t *testing.T) {
	url, err := GetRandomDogOfSubBreed("hound", "afghan")
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(url, "http"))
}

func TestGetBunchOfRandomDogsOfSubBreed(t *testing.T) {
	imageCount := 50
	urls, err := GetBunchOfRandomDogsOfSubBreed("hound", "afghan", imageCount)
	assert.Nil(t, err)
	for _, url := range urls {
		assert.True(t, strings.HasPrefix(url, "http"))
	}
	assert.Equal(t, len(urls), imageCount)
}
