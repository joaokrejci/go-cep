package control

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joaokrejci/go-cep/entities"
	"github.com/spf13/viper"
)

type TestCase struct {
	input interface{}
	expected interface{}
	behavior string
}

var cases []TestCase = []TestCase{
	{
		behavior: "Should return when cep is not formatted",
		input: "01001000", 
		expected: entities.Address{
		Street: "Praça da Sé",
		Complement: "lado ímpar",
		Neighbourhood: "Sé",
		City: "São Paulo",
		Uf: "SP",
		Cep: "01001-000",
	}},
	{
		behavior: "Should return when cep has special characters",
		input: "01001--%ab ===000", 
		expected: entities.Address{
		Street: "Praça da Sé",
		Complement: "lado ímpar",
		Neighbourhood: "Sé",
		City: "São Paulo",
		Uf: "SP",
		Cep: "01001-000",
	}},
	{
		behavior: "Should return nothing when cep has wrong number of characters",
		input: "010010000",  // 9
		expected: entities.Address{},
	},
}

func TestGetCep(t *testing.T) {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()
	if err != nil {
		t.Errorf("Erro ao ler o arquivo de configuração %s", err)
	}

	for _, testCase := range cases {
		address, err := GetCep(testCase.input.(string))
		if err != nil {
			t.Errorf(err.Error())
		}

		if !cmp.Equal(address, testCase.expected) {
			t.Errorf("%v\n GetCep failed for case %v. Expected %v, got %v", testCase.behavior, testCase.input, testCase.expected, address)
		}
	}
}