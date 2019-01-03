package main

import (
	"fmt"
)

const (
	h1 = "Endereço que deseja enviar requisições; 80 para HTTP ou 443 para HTTPS"
)

func help() string {
	u := fmt.Sprintf("Veja" + h1)

	return u
}
