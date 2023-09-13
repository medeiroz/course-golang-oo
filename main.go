package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaFlavio := ContaCorrente{
		titular:       "Flavio",
		numeroAgencia: 54,
		numeroConta:   123456,
		saldo:         75.45,
	}

	fmt.Println("contaFlavio", contaFlavio)

	var contaBia *ContaCorrente
	contaBia = new(ContaCorrente)
	contaBia.titular = "Bia"
	contaBia.numeroAgencia = 87
	contaBia.numeroConta = 6543211
	contaBia.saldo = 125.40

	fmt.Println("contaBia", *contaBia)
}
