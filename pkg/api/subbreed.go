package api

func GetAllBreedsWithSubBreeds() (map[string][]string, error) {
	var data struct {
		Breeds map[string][]string `json:"message"`
	}
	err := unmarshalFromURL("https://dog.ceo/api/breeds/list/all", &data)

	if err != nil {
		return nil, err
	}

	return data.Breeds, nil
}
