package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaFlavio := contas.ContaCorrente{
		Titular: "Flavio",
		Saldo:   75.45,
	}

	contaBia := contas.ContaCorrente{
		Titular: "Bia",
		Saldo:   456,
	}

	fmt.Println("contaFlavio", contaFlavio)
	fmt.Println("contaBia", contaBia)

	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, -50.0))
}
