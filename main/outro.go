package main

import "fmt"

// teste func publica
func Outro() {
	fmt.Println("Teste outro, publico, mesmo diretorio")
	outro2()
}

// teste func privada, não exibe
func outro2() {
	fmt.Println("Teste outro 2, privado, mesmo diretorio, uso interno")
}
