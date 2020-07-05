package main 

import (
	"fmt"
	"sync"
)

type ChopS struct { sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	// fmt.Println("hello")

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new (ChopS)
	}

	philos := make([]*Philo, 5)

	for i:= 0; i<5; i++ {
		philos[i] = &Philo{CSticks[i],CSticks[(i+1)%5]}
	}

	for i:= 0; i<5; i++ {
		philos[i].eat()
	}

}