package main

import (
	"fmt"
	"strconv"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor <= conta.saldo

	if podeSacar {
		conta.saldo -= valor
		return "Saque realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 0, 64)
	}

	return "Saldo insulficiente"
}

func (conta *ContaCorrente) Depositar(valor float64) string {
	podeDepositar := valor > 0

	if podeDepositar {
		conta.saldo += valor
	}

	return "Deposito realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 0, 64)
}

func main() {

	contaFlavio := ContaCorrente{
		titular:       "Flavio",
		numeroAgencia: 54,
		numeroConta:   123456,
		saldo:         75.45,
	}

	fmt.Println("Saldo antes do saque", contaFlavio.saldo)

	fmt.Println("Saque R$ 50,00", contaFlavio.Sacar(50.0))
	fmt.Println("Saque R$ 20,00", contaFlavio.Sacar(20.0))
	fmt.Println("Saque R$ 4,00", contaFlavio.Sacar(4.0))
	fmt.Println("Saque R$ 2,00", contaFlavio.Sacar(2.0))
	fmt.Println("Depositar R$ 50,00", contaFlavio.Depositar(50.0))
	fmt.Println("Saque R$ 10,00", contaFlavio.Sacar(10.0))
}
