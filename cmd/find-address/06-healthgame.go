/**
 * Veja que nesse caso a função é a mesma, 
 * porém agora estamos usando um ponteiro intermediario para nao precisar chamar a funcao **hitPlayer** duas vezes
 * isso faz com que nosso aob anterior nao funcione mais como esperado, agora ele congela o health de todos os players
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
		var p *int
		switch n {
		case 1:
			p = &p1
		case 2:
			p = &p2
		default :
			log.Printf("unknow player %v", s)
			continue;
		}
		hitPlayer(p)
		
	}
}

func createPlayer() int {
	health := 100
	return health
}
func hitPlayer(health *int){
	*health -= 1;
}