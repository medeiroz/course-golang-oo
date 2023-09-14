package contas

import (
	"banco/clientes"
	"strconv"
)

type ContaPoupanca struct {
	Titular  clientes.Titular
	Operacao int
	saldo    float64
}

func (conta *ContaPoupanca) Sacar(valor float64) string {
	if valor > conta.saldo {
		return "Falha ao realizar saque. saldo insulficiente"
	}

	conta.saldo -= valor
	return "Saque realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 2, 64)
}

func (conta *ContaPoupanca) Depositar(valor float64) string {
	if valor < 0 {
		return "Falha ao depositar. Não é permitido depositar valores negativos"
	}

	conta.saldo += valor
	return "Deposito realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 2, 64)
}

func (conta *ContaPoupanca) ObterSaldo() float64 {
	return conta.saldo
}
