package contas

import (
	"banco/clientes"
	"strconv"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Saldo   float64
}

func (conta *ContaCorrente) Sacar(valor float64) string {
	if valor > conta.Saldo {
		return "Falha ao realizar saque. Saldo insulficiente"
	}

	conta.Saldo -= valor
	return "Saque realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.Saldo, 'f', 2, 64)
}

func (conta *ContaCorrente) Depositar(valor float64) string {
	if valor < 0 {
		return "Falha ao depositar. Não é permitido depositar valores negativos"
	}

	conta.Saldo += valor
	return "Deposito realizado com sucesso. Novo saldo: " + strconv.FormatFloat(conta.Saldo, 'f', 2, 64)
}

func (conta *ContaCorrente) Transferir(contaDestino *ContaCorrente, valor float64) string {
	if valor < 0 {
		return "Não é permitido transferir valores negativos"
	}

	if valor > conta.Saldo {
		return "Falha ao realizar saque. Saldo insulficiente"
	}

	conta.Saldo -= valor
	contaDestino.Depositar(valor)

	return "Transferencia realizada com sucesso. Novo saldo: " + strconv.FormatFloat(conta.Saldo, 'f', 2, 64)
}
