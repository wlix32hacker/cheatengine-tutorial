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

type Player struct {
	Health int
	Type PlayerType
}

type PlayerType int;

const(
	PLAYER = 1
	ENEMY = 2
)

func main(){
	p1, p2 := createPlayer(PLAYER), createPlayer(ENEMY);
	in := bufio.NewReader(os.Stdin)
	for ; ; {
		log.Printf("p1=%d, p2=%d", p1, p2)
		l, _, _ := in.ReadLine()
		s := string(l)
		n, _ := strconv.Atoi(s)
		var p *Player
		switch n {
		case 1:
			p = p1
		case 2:
			p = p2
		default :
			log.Printf("unknow player %v", s)
			continue;
		}
		hitPlayer(p)
	}
}

func createPlayer(ptype PlayerType) *Player {
	return &Player{100, ptype}
}

func hitPlayer(p *Player){
	p.Health -= 1;
}