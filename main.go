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

	contaPoupancaJoao := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "Joao",
			CPF:       "999.999.999-99",
			Profissao: "PO",
		},
	}

	contaFlavio.Depositar(100)
	contaBia.Depositar(654)
	contaPoupancaJoao.Depositar(40)

	fmt.Println("contaFlavio", contaFlavio, contaFlavio.ObterSaldo())
	fmt.Println("contaBia", contaBia, contaBia.ObterSaldo())
	fmt.Println("contaPoupancaJoao", contaPoupancaJoao, contaPoupancaJoao.ObterSaldo())

	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, -50.0))
}
