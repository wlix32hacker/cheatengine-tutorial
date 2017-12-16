package main 

import (
	"bufio"
	"os"
	"log"
	"strconv"
)

func main(){

	var health ***int = new(**int)
	var n int = 0

	*health = new(*int)
	**health = new(int)
	***health = 1000

	log.Printf("Press the value to decrease and type enter, type 4 to change the health pointer")
	in := bufio.NewReader(os.Stdin)
	for ; ; {
		
		l, _, _ := in.ReadLine()
		n, _ = strconv.Atoi(string(l))

		tmp := ***health

		if n == 4 {
			health = new(**int)
			*health = new(*int)
			**health = new(int)
		}
		
		***health = tmp - n
		log.Printf("health=%p, *health=%p, **health=%p, ***health=%d, &health=%p", health, *health, **health, ***health, &health)
		log.Print()

	}
}

