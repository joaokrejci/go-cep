package entities

import "fmt"

type Address struct {
	Street string `json:"logradouro"`
	Uf 	string	`json:"uf"`
	City string `json:"localidade"`
	State string `json:"estado"`
	Complement string `json:"complemento"`
	Neighbourhood string `json:"bairro"`
	Cep string `json:"cep"`
}

var ufs = map[string]string {
	"Acre": "AC",
	"Alagoas": "AL",
	"Amapá": "AP",
	"Amazonas": "AM",
	"Bahia": "BA",
	"Ceará": "CE",
	"Distrito Federal": "DF",
	"Espírito Santo": "ES",
	"Goiás": "GO",
	"Maranhão": "MA",
	"Mato Grosso": "MT",
	"Mato Grosso do Sul": "MS",
	"Minas Gerais": "MG",
	"Pará": "PA",
	"Paraíba": "PB",
	"Paraná": "PR",
	"Pernambuco": "PE",
	"Piauí": "PI",
	"Rio de Janeiro": "RJ",
	"Rio Grande do Norte": "RN",
	"Rio Grande do Sul": "RS",
	"Rondônia": "RO",
	"Roraima": "RR",
	"Santa Catarina": "SC",
	"São Paulo": "SP",
	"Sergipe": "SE",
	"Tocantins": "TO",
}

func (address *Address) IsRequiredFieldsPresent() bool {
	return address.City != "" && address.Street != "" && (address.Uf != "" || address.State != "")
}

func (address *Address) Assert() error {
	if !address.IsRequiredFieldsPresent() {
		return fmt.Errorf("o comando busca exige o uso das flags -l, -C, -u e/ou -e")
	}

	if(address.Uf == "") {
		address.Uf = convertStateToUf(address.State)
		if(address.Uf == "") {
			return fmt.Errorf("estado %v inválido", address.State)
		}
	}

	return nil
}

func convertStateToUf(state string) string {
	return ufs[state]
}