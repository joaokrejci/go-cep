package control

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joaokrejci/go-cep/entities"
	"github.com/spf13/viper"
)

var testCases []TestCase = []TestCase{
	{
		behavior: "Should get the right address based on search (not cep)",
		input: entities.Address{
			Street: "Rua São João Da Barra",
			City: "Rio de Janeiro",
			Uf: "RJ",
		},
		expected: []entities.Address{
			{
				Cep: "23026-350",
				Street: "Rua São João da Barra",
				Neighbourhood: "Pedra de Guaratiba",
				City: "Rio de Janeiro",
				Uf: "RJ",
			},
		},
	},
	{
		behavior: "Should get the right address based on search (not cep) (asserting address)",
		input: entities.Address{
			Street: "Rua São João Da Barra",
			City: "Rio de Janeiro",
			State: "Rio de Janeiro",
		},
		expected: []entities.Address{
			{
				Cep: "23026-350",
				Street: "Rua São João da Barra",
				Neighbourhood: "Pedra de Guaratiba",
				City: "Rio de Janeiro",
				Uf: "RJ",
			},
		},
	},
	{
		behavior: "Should error if not all address information is present",
		input: entities.Address{
			Street: "Rua São João Da Barra",
			City: "Rio de Janeiro",
		},
		expected: fmt.Errorf("o comando busca exige o uso das flags -l, -C, -u e/ou -e"),
	},
	{
		behavior: "Should error if state is invalid",
		input: entities.Address{
			Street: "Rua São João Da Barra",
			City: "Rio de Janeiro",
			State: "California",
		},
		expected: fmt.Errorf("estado %v inválido", "California"),
	},
}

func TestSearchByAddress(t * testing.T) {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()
	if err != nil {
		t.Errorf("Erro ao ler o arquivo de configuração %s", err)
	}

	for _, testCase := range testCases {
		address := testCase.input.(entities.Address)
		output, err := SearchByAddress(&address)
		if err != nil {
			expectedError := testCase.expected.(error)
			if err.Error() != expectedError.Error() {
				t.Errorf("Unexpected error. Expected %v, got %v", expectedError.Error(), err.Error())
			}
			continue
		}

		if !cmp.Equal(output, testCase.expected) {
			t.Errorf("%v\n GetCep failed for case %v. Expected %v, got %v", testCase.behavior, testCase.input, testCase.expected, output)
		}
	}
	
}