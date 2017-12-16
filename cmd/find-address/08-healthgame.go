/**
 * Fiz esse exemplo tentando simular o exemplo 8 para que o AOB nao funcione quando trocar o ponteiro mas continua funcionando
 * portanto esse exemplo tem o mesmo efeito do 07
 */
package main 

import (
	"bufio"
	"os"
	"log"
	"strconv"
	"math/rand"
	"C"
	"unsafe"
)

type Player4 struct {
	a int
	b int
	c int
	d int
	e int
	f int
	health int
}

type Player3 struct {
	p *Player4
	a int
	b int
	c int
	d int
	e int
	f int
}

type Player2 struct {
	a int
	b int
	c int
	d int
	e int
	p *Player3
	f int
}

type Player1 struct {
	a int
	b int
	p *Player2
	c int
	d int
	e int
	f int
}

func main(){

	log.Printf("%d", unsafe.Sizeof(*new(Player1)))
	var l1 *Player1 = new(Player1)
	// l1 = (*Player1)( C.malloc( C.size_t(  + uintptr(rand.Intn(128)) ) ) ) 
	realloc(l1)

	log.Printf("Press the value to decrease and type enter, type 4 to change the health pointer")
	in := bufio.NewReader(os.Stdin)
	for ; ; {
		
		l, _, _ := in.ReadLine()
		n, _ := strconv.Atoi(string(l))

		if n == 4 {
			l1 = new(Player1)
//			l1 = (*Player1)( C.malloc( C.size_t( unsafe.Sizeof(*new(Player1)) + uintptr(rand.Intn(128)) ) ) )
			realloc(l1)
		}
		(*l1.p.p.p).health -= n

		log.Printf("l1 address=%p, p=%p, p-value=%#v", &l1, l1, l1)
		log.Printf("l2 address=%p, p=%p, p-value=%#v", &l1.p, l1.p, l1.p)
		log.Printf("l3 address=%p, p=%p, p-value=%#v", &l1.p.p, l1.p.p, l1.p.p)
		log.Printf("l4 address=%p, p=%p, p-value=%#v", &l1.p.p.p, l1.p.p.p, l1.p.p.p)
		log.Print()

	}
}

func realloc(l1 *Player1){
	l1.a = rand.Intn(4000)
	l1.b = rand.Intn(4000)
	l1.c = rand.Intn(4000)
	l1.d = rand.Intn(4000)
	l1.e = rand.Intn(4000)
	l1.f = rand.Intn(4000)
	 l1.p = new(Player2)
//	l1.p = (*Player2)( C.malloc( C.size_t( unsafe.Sizeof(*new(Player2)) + uintptr(rand.Intn(128)) ) ) )

	l1.p.a = rand.Intn(4000)
	l1.p.b = rand.Intn(4000)
	l1.p.c = rand.Intn(4000)
	l1.p.d = rand.Intn(4000)
	l1.p.e = rand.Intn(4000)
	l1.p.f = rand.Intn(4000)
	l1.p.p = new(Player3)
//	l1.p.p = (*Player3)( C.malloc( C.size_t( unsafe.Sizeof(*new(Player3)) + uintptr(rand.Intn(128)) ) ) )

	
	l1.p.p.a = rand.Intn(4000)
	l1.p.p.b = rand.Intn(4000)
	l1.p.p.c = rand.Intn(4000)
	l1.p.p.d = rand.Intn(4000)
	l1.p.p.e = rand.Intn(4000)
	l1.p.p.f = rand.Intn(4000)
	l1.p.p.p = new(Player4)
//	l1.p.p.p = (*Player4)( C.malloc( C.size_t( unsafe.Sizeof(*new(Player4)) + uintptr(rand.Intn(128)) ) ) )

	l1.p.p.p.a = rand.Intn(4000)
	l1.p.p.p.b = rand.Intn(4000)
	l1.p.p.p.c = rand.Intn(4000)
	l1.p.p.p.d = rand.Intn(4000)
	l1.p.p.p.e = rand.Intn(4000)
	l1.p.p.p.f = rand.Intn(4000)
	l1.p.p.p.health = rand.Intn(4000)
}