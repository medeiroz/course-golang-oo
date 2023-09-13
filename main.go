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
	}

	contaBia := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bia",
			CPF:       "111.222.333-44",
			Profissao: "QA",
		},
	}

	contaFlavio.Depositar(100)
	contaBia.Depositar(654)

	fmt.Println("contaFlavio", contaFlavio, contaFlavio.ObterSaldo())
	fmt.Println("contaBia", contaBia, contaBia.ObterSaldo())

	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, -50.0))
}
