/**
 * Nesse cenario ja muda de figura
 * A funcao createPlayer é a responsavel por retornar o valor e cada vez que ela for chamada será retornada uma instancia
 * por isso temos um endereco dinamico aqui
 * 
 * De qualquer forma nesse caso basta procurar quem escreve e fazer um AOB e tudo vai funcionar,
 * alem disso nao vai ter o problema de todo mundo ter o sangue locado,
 * apenas um nao vai perder o sangue
 * 
 */
package main 

import (
	"bufio"
	"os"
	"log"
	"strconv"
)

func main(){
	p1, p2 := createPlayer(), createPlayer();
	in := bufio.NewReader(os.Stdin)
	for ; ; {
		log.Printf("p1=%d, p2=%d", p1, p2)
		l, _, _ := in.ReadLine()
		s := string(l)
		n, _ := strconv.Atoi(s)
		switch n {
		case 1:
			p1 -= 1
		case 2:
			p2 -= 1
		default :
			log.Printf("unknow player %v", s)
			continue;
		}
		
	}
}

func createPlayer() int {
	health := 100
	return health
}