/**
 * Mesmo instanciando a variavel dentro da funcao ela ainda é estática pois é declarada
 * e nao criada com uma funcao como o new (malloc)
 */
package main 

import (
	"bufio"
	"os"
	"log"
)


func main(){
	health := 100
	in := bufio.NewReader(os.Stdin)
	for ; ; {
		log.Printf("health=%d", health)
		_, _, _ = in.ReadLine()
		health -= 1
	}
}