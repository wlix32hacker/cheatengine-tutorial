/**
 * Aqui foi criada uma funcao compartilhada para tirar o sangue, 
 * mas isso nao muda em nada o assembly e o mesmo aob do anterior continua funcionando
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
			hitPlayer(&p1)
		case 2:
			hitPlayer(&p2)
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
func hitPlayer(health *int){
	*health -= 1;
}