package main

import (
	"fmt"
	"math/rand"
	"time"
)

// VARIABLES TO EXPERIMENT
const (
	size              int  = 2
	incomingInterval  int  = 1
	incomingRandom    bool = true
	outcomingInterval int  = 10
	outcomingRandom   bool = false
)

type Balancer struct {
	incoming  chan int
	available chan int
	size      int
}

func CreateBalancer(n int) Balancer {
	var b Balancer = Balancer{
		incoming:  make(chan int),
		available: make(chan int, n),
		size:      n,
	}
	for i := 0; i < n; i++ {
		b.available <- i
	}
	return b
}

func (b *Balancer) runIncoming() {
	if incomingRandom == true {
		for {
			time.Sleep(time.Second * time.Duration(rand.Intn(incomingInterval)+1))
			b.incoming <- 0
		}
	} else {
		for {
			time.Sleep(time.Second * time.Duration(incomingInterval))
			b.incoming <- 0
		}
	}
}

func (b *Balancer) outcoming(n int, position int) {
	var timeToWait time.Duration

	if outcomingRandom == true {
		timeToWait = (time.Second * time.Duration(rand.Intn(outcomingInterval)+1))
	} else {
		timeToWait = time.Second * time.Duration(outcomingInterval)
	}

	time.Sleep(timeToWait)
	b.available <- position
	fmt.Printf("unit number %d frees resource %d\n", n, position)
}

func (b *Balancer) run() {
	for n := 0; ; n++ {
		<-b.incoming
		var position int = <-b.available
		fmt.Printf("unit number %d takes resource %d\n", n, position)
		go b.outcoming(n, position)
	}
}

func main() {
	var b Balancer = CreateBalancer(size)
	go b.runIncoming()
	b.run()
}
