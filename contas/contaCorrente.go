package contas

import (
	"banco/clientes"
	"strconv"
)

type ContaCorrente struct {
	Titular clientes.Titular
	saldo   float64
}

func (conta *ContaCorrente) Sacar(valor float64) string {
	if valor > conta.saldo {
		return "Falha ao realizar saque. saldo insulficiente"
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
		return "Falha ao realizar saque. saldo insulficiente"
	}

	conta.saldo -= valor
	contaDestino.Depositar(valor)

	return "Transferencia realizada com sucesso. Novo saldo: " + strconv.FormatFloat(conta.saldo, 'f', 2, 64)
}

func (conta *ContaCorrente) ObterSaldo() float64 {
	return conta.saldo
}
