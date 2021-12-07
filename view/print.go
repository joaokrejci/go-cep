package view

import (
	"fmt"

	"github.com/joaokrejci/go-cep/entities"
)

func PrintAddress(address entities.Address) {
	formatString := fmt.Sprintf("%s: %s\n%s: %s\n%s: %s\n%s: %s\n%s: %s\n", 
	Teal("Logradouro"), Yellow(address.Street),
	Teal("Complemento"), Yellow(address.Complement),
	Teal("Bairro"), Yellow(address.Neighbourhood),
	Teal("Localidade"), Yellow(address.City),
	Teal("UF"), Yellow(address.Uf))
	
	fmt.Print(formatString)
}

func PrintMultipleAddresses(addresses []entities.Address) {
	formatString := fmt.Sprintf("\n%s. %s \n\t%s", "%d", Teal("%s"), Yellow("%v - %v - %v - %v"))
	for index, value := range addresses {
		fmt.Printf(formatString, index + 1, value.Cep, value.Street, value.Neighbourhood, value.City, value.Uf)
	}
}