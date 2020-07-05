package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const PHILO_NUM int = 5
const STICKS_NUM int = 5
const PHILO_MAX_EATING int = 2
const PHILO_MAX_EATS int = 3

type Host struct {
	p []*Philo
	s []*ChopS
	c chan int
}

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func (p *Philo) selectSticks(MAX int) (int, int) {
	left := rand.Intn(MAX)
	right := left
	for left == right {
		right = rand.Intn(MAX)
	}
	return left, right
}

func (p *Philo) eat(c1 chan int, sa []*ChopS) {

	l, r := p.selectSticks(STICKS_NUM)
	p.leftCS, p.rightCS = sa[l], sa[r]
	for i := 0; i < PHILO_MAX_EATS; i++ {

		<-c1 // can eat?

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %v\n", p.id)

		fmt.Printf("finishing eating %v\n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func (h *Host) Setup() {
	h.s = StickFactory(STICKS_NUM)
	h.p = PhiloFactory(PHILO_NUM)
	h.c = make(chan int, PHILO_MAX_EATING)
}

func (h *Host) Start() {
	for i := 0; i < PHILO_NUM; i++ {
		go h.p[i].eat(h.c, h.s)
	}
}

func (h *Host) Wait() {
	max_exec := len(h.p) * PHILO_MAX_EATS
	for i := 0; i < max_exec; i++ {
		h.c <- 1
	}
}

func StickFactory(num int) []*ChopS {
	s := make([]*ChopS, num)
	for i := 0; i < num; i++ {
		s[i] = new(ChopS)
	}
	return s
}

func PhiloFactory(num int) []*Philo {
	p := make([]*Philo, num)
	for i := 0; i < num; i++ {
		p[i] = &Philo{}
		p[i].id = i
	}
	return p
}

func main() {
	// 1 seed random
	rand.Seed(time.Now().UnixNano())

	h := Host{}
	h.Setup()

	h.Start()
	h.Wait()

}
