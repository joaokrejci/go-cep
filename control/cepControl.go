package control

import (
	"encoding/json"
	"fmt"
	"path"
	"regexp"

	"github.com/joaokrejci/go-cep/api"
	"github.com/joaokrejci/go-cep/entities"
	"github.com/spf13/viper"
)



func formatCep(cep string) (formattedCep string, err error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return
	}

	formattedCep = reg.ReplaceAllString(cep, "")
	return
}

// GetCep retrieves an address information based on a given cep
func GetCep(cep string) (address entities.Address, err error) {
	cleanCep, err := formatCep(cep)
	if err != nil {
		return
	}

	apiAddress, ok := viper.Get("API_ADDRESS").(string)
	if !ok {
		err = fmt.Errorf("erro ao obter configuração API_ADDRESS. Verifique o arquivo config.yml")
		return
	}

	url := apiAddress + path.Join(cleanCep, "json")
	body, err := api.Request(url) 
	if err != nil {
		return
	}

	json.Unmarshal(body, &address)

	return
}