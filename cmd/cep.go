/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/joaokrejci/go-cep/control"
	"github.com/joaokrejci/go-cep/view"
	"github.com/spf13/cobra"
)

// cepCmd represents the cep command
var cepCmd = &cobra.Command{
	Use:   "cep",
	Short: "Busca um endereço baseado em um cep",
	Long: `
Busca um endereço baseado em um cep
Utilização:
	$ go-cep cep 01001-000

	Logradouro: Praça da Sé
	Complemento: Lado Ímpar
	Bairro: Sé
	Localidade: São Paulo
	UF: SP
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(cmd.Long)
			return
		}

		address, err := control.GetCep(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		view.PrintAddress(address)
	},
}

func init() {
	rootCmd.AddCommand(cepCmd)
}
