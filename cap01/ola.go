/*
# Arquivos-fonte devem seguir sempre a mesma estrutura, dividida em três seções:
*/

/*
# Declaração do pacote

# Todo código Go deve obrigatoriamente existir dentro de um pacote

# Todo programa em Go deve ter um pacote 'main' contendo um função main()
*/
package main

/*
# Declaraçao dos pacotes externos dos quais o arquivo-fonte depende
*/
import "fmt"

/*
# Código referente ao programa ou pacote
# A função main() não recebe nenhum argumento e não retorna nenhum valor
*/
func main() {
	fmt.Println("Olá, Go!")
}
