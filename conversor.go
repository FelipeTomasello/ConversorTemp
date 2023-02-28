package main

import "fmt"

const ebulicaoAguaK = 373

func main() {
	ebulicaoAguaC := ebulicaoAguaK - 273

	fmt.Printf("A temperatura da água em K é: %v, e a temperatura da água em C é: %v", ebulicaoAguaK, ebulicaoAguaC)
}
