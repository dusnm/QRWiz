package countries

import (
	"encoding/json"

	"github.com/dusnm/QRWiz/pkg/assets"
)

type (
	Country struct {
		Name string `json:"name"`
		CCA2 string `json:"cca2"`
		CCA3 string `json:"cca3"`
	}
)

func All() ([]Country, error) {
	countries := make([]Country, 0, 200)
	list := assets.ListOfCountries()

	if err := json.Unmarshal(list.StaticContent, &countries); err != nil {
		return nil, err
	}

	return countries, nil
}

func AllIndexedByName() (map[string]Country, error) {
	list, err := All()
	if err != nil {
		return nil, err
	}

	countries := make(map[string]Country, len(list))
	for _, country := range list {
		countries[country.Name] = country
	}

	return countries, nil
}
