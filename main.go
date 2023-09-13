package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	contaFlavio := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Flavio",
			CPF:       "123.123.123-12",
			Profissao: "Desenvolvedor",
		},
		Saldo: 75.45,
	}

	contaBia := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bia",
			CPF:       "111.222.333-44",
			Profissao: "QA",
		},
		Saldo: 456,
	}

	fmt.Println("contaFlavio", contaFlavio)
	fmt.Println("contaBia", contaBia)

	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, -50.0))
}
