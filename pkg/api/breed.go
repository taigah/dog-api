package api

func GetAllBreeds() ([]string, error) {
	var data struct {
		Breeds map[string][]string `json:"message"`
	}
	err := unmarshalFromURL("https://dog.ceo/api/breeds/list/all", &data)

	if err != nil {
		return nil, err
	}

	breeds := make([]string, 0)

	for breed, _ := range data.Breeds {
		breeds = append(breeds, breed)
	}

	return breeds, nil
}
