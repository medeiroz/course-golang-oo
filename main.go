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
		numeroAgencia: 0054,
		numeroConta:   123456,
		saldo:         75.45,
	}

	fmt.Println(contaFlavio)
}
