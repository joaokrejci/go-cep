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
	"log"

	"github.com/joaokrejci/go-cep/control"
	"github.com/joaokrejci/go-cep/entities"
	"github.com/joaokrejci/go-cep/view"
	"github.com/spf13/cobra"
)


var Address entities.Address

// buscaCmd represents the busca command
var buscaCmd = &cobra.Command{
	Use:   "busca",
	Short: "Busca CEPs baseado no endereço completo",
	Long: `
Busca CEPs baseado no endereço completo
Utilização:
	$ go-cep busca -l "Afonso Pena" -e "Mato Grosso do Sul" -C "Campo Grande"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		addresses, err := control.SearchByAddress(&Address)
		if err != nil {
			log.Fatal(err.Error())
		}

		view.PrintMultipleAddresses(addresses)
	},
}

func init() {
	rootCmd.AddCommand(buscaCmd)

	buscaCmd.PersistentFlags().StringVarP(&Address.Street, "logradouro", "l", "", "Especifica um logradouro para busca textual")
	buscaCmd.PersistentFlags().StringVarP(&Address.Uf, "uf", "u", "", "Especifica uma unidade federativa para busca textual")
	buscaCmd.PersistentFlags().StringVarP(&Address.City, "cidade", "C", "", "Especifica uma cidade para busca textual")
	buscaCmd.PersistentFlags().StringVarP(&Address.State, "estado", "e", "", "Especifica um estado para busca textual")
}
