package control

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/joaokrejci/go-cep/api"
	"github.com/joaokrejci/go-cep/entities"
	"github.com/spf13/viper"
)


func SearchByAddress(address *entities.Address) (addresses []entities.Address, err error) {
	err = address.Assert()
	if err != nil {
		return
	}

	apiAddress, ok := viper.Get("API_ADDRESS").(string)
	if !ok {
		err = fmt.Errorf("erro ao obter configuração API_ADDRESS")
		return
	}

	url := apiAddress + path.Join(address.Uf, address.City, address.Street, "json")
	body, err := api.Request(url)
	if err != nil {
		return
	}

	json.Unmarshal(body, &addresses)

	return
}