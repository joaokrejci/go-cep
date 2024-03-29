/*
Copyright © 2021 joaokrejci <joao.krejci@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cep",
	Short: "Buscas por CEP na linha de comando",
	Long: `	
Encontre informações de endereço a partir de um CEP ou procure por um CEP a partir do nome de uma rua

> Comandos

	'cep' busca padrão por cep
		Especifica uma busca por cep (pode ser omitido)

	'busca' busca por endereço
		Utilizada para realizar a busca através do endereço em lugar da busca padrão por CEP
		Deve-se utilizar em conjunto com as opções '-l','-e', '-u' e '-C'

>Flags

	-l ou --logradouro
		Especifica um logradouro para busca textual

	-u ou --uf
		Especifica a Unidade Federativa ou Estado a ser buscado

	-e ou --estado
		Especifica um estado, pode ser utilizada no lugar de --uf especificando o nome de um estado em vez de uma unidade federativa.

	-C ou --cidade
		Especifica uma cidade a ser buscada

Exemplos
	$ go-cep cep 01001-000

	Logradouro: Praça da Sé
	Complemento: Lado Ímpar
	Bairro: Sé
	Localidade: São Paulo
	UF: SP

	$ go-cep busca -l "São João" -u "SP" -C "São Paulo"

	1. 08245-040
		Avenida São José do Cedro - Vila Progresso (Zona Leste) - São Paulo - SP
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
