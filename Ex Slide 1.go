package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase1 := scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	frase2 := scanner.Text()

	s1 := frase1 + " " + frase2

	fmt.Print("As frases concatenadas são: ", s1)
}
