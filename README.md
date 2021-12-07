# go-cep

Encontre informações de endereço a partir de um CEP ou procure por um CEP a partir do nome de uma rua

## Utilização

``` shell
go-cep 01001-000

Logradouro: Praça da Sé
Complemento: Lado Ímpar
Bairro: Sé
Localidade: São Paulo
UF: SP

```

``` shell
go-cep -b -l "São João" -u "SP" -C "São Paulo"

1. 08245-040
   Avenida São José do Cedro - Vila Progresso (Zona Leste) - São Paulo - SP
```

## Comandos

### `cep` busca padrão por cep

Especifica uma busca por cep (pode ser omitido)

### `busca` busca por endereço

Utilizada para realizar a busca através do endereço em lugar da busca padrão por CEP
Deve-se utilizar em conjunto com as opções `-l`,`-e`, `-u` e `-C`

#### Flags

##### `-l` ou `--logradouro`

Especifica um logradouro para busca textual

##### `-u` ou `--uf`

Especifica a Unidade Federativa ou Estado a ser buscado

##### `-e` ou `--estado`

Especifica um estado, pode ser utilizada no lugar de `--uf` especificando o nome de um estado em vez de uma unidade federativa.

##### `-C` ou `--cidade`

Especifica uma cidade a ser buscada

<small>Esta aplicação utiliza a API [ViaCEP](https://viacep.com.br/)</small>
