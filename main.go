package main

import (
	"fmt"
	"strconv"
)

type ContaCorrente struct {
	titular string
	saldo   float64
}

func (conta *ContaCorrente) Sacar(valor float64) string {
	if valor > conta.saldo {
		return "Falha ao realizar saque. Saldo insulficiente"
	}

	conta.saldo -= valor
	return "Saque realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 2, 64)
}

func (conta *ContaCorrente) Depositar(valor float64) string {
	if valor < 0 {
		return "Falha ao depositar. Não é permitido depositar valores negativos"
	}

	conta.saldo += valor
	return "Deposito realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 2, 64)
}

func (conta *ContaCorrente) Transferir(contaDestino *ContaCorrente, valor float64) string {
	if valor < 0 {
		return "Não é permitido transferir valores negativos"
	}

	if valor > conta.saldo {
		return "Falha ao realizar saque. Saldo insulficiente"
	}

	conta.saldo -= valor
	contaDestino.Depositar(valor)

	return "Transferencia realizada com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 2, 64)
}

func main() {

	contaFlavio := ContaCorrente{
		titular: "Flavio",
		saldo:   75.45,
	}

	contaBia := ContaCorrente{
		titular: "Bia",
		saldo:   456,
	}

	fmt.Println("contaFlavio", contaFlavio)
	fmt.Println("contaBia", contaBia)

	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, 50.0))
	fmt.Println(contaFlavio.Transferir(&contaBia, -50.0))
}
