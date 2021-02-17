package api

import "fmt"

func GetRandomDogImageUrl() (string, error) {
	var data struct {
		URL string `json:"message"`
	}
	err := unmarshalFromURL("https://dog.ceo/api/breeds/image/random", &data)

	if err != nil {
		return "", err
	}

	return data.URL, nil
}

func GetBunchOfRandomDogImagesUrl(imageCount int) ([]string, error) {
	var data struct {
		URLs []string `json:"message"`
	}

	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breeds/image/random/%v", imageCount), &data)
	if err != nil {
		return nil, err
	}

	return data.URLs, nil
}

func GetAllDogsOfBreed(breed string) ([]string, error) {
	var data struct {
		URLs []string `json:"message"`
	}
	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/images", breed), &data)

	if err != nil {
		return nil, err
	}

	return data.URLs, nil
}

func GetRandomDogOfBreed(breed string) (string, error) {
	var data struct {
		URL string `json:"message"`
	}
	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/images/random", breed), &data)

	if err != nil {
		return "", err
	}

	return data.URL, nil
}

func GetBunchOfRandomDogsOfBreed(breed string, imageCount int) ([]string, error) {
	var data struct {
		URLs []string `json:"message"`
	}

	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/images/random/%v", breed, imageCount), &data)
	if err != nil {
		return nil, err
	}

	return data.URLs, nil
}

func GetSubBreeds(breed string) ([]string, error) {
	var data struct {
		SubBreeds []string `json:"message"`
	}

	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/list", breed), &data)
	if err != nil {
		return nil, err
	}

	return data.SubBreeds, nil
}

func GetDogsOfSubBreed(breed string, subBreed string) ([]string, error) {
	var data struct {
		URLs []string `json:"message"`
	}
	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/%v/images", breed, subBreed), &data)
	if err != nil {
		return nil, err
	}
	return data.URLs, nil
}

func GetRandomDogOfSubBreed(breed string, subBreed string) (string, error) {
	var data struct {
		URL string `json:"message"`
	}
	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/%v/images/random", breed, subBreed), &data)
	if err != nil {
		return "", err
	}
	return data.URL, nil
}

func GetBunchOfRandomDogsOfSubBreed(breed string, subBreed string, imageCount int) ([]string, error) {
	var data struct {
		URLs []string `json:"message"`
	}
	err := unmarshalFromURL(fmt.Sprintf("https://dog.ceo/api/breed/%v/%v/images/random/%v", breed, subBreed, imageCount), &data)
	if err != nil {
		return nil, err
	}
	return data.URLs, nil
}