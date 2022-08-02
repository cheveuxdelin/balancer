package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gookit/color"
)

// VARIABLES TO EXPERIMENT
const (
	size              int  = 10
	incomingInterval  int  = 1
	incomingRandom    bool = false
	outcomingInterval int  = 15
	outcomingRandom   bool = true
)

type Balancer struct {
	resources []bool
	incoming  chan bool
	available chan int
}

func CreateBalancer(n int) Balancer {
	var b Balancer = Balancer{
		resources: make([]bool, n),
		incoming:  make(chan bool),
		available: make(chan int, n),
	}
	for i := 0; i < n; i++ {
		b.available <- i
	}
	return b
}

func (b *Balancer) runIncoming() {
	if incomingRandom {
		for {
			time.Sleep(time.Second * time.Duration(rand.Intn(incomingInterval)+1))
			b.incoming <- true
		}
	} else {
		for {
			time.Sleep(time.Second * time.Duration(incomingInterval))
			b.incoming <- true
		}
	}
}

func (b *Balancer) outcoming(n int, position int) {
	var timeToWait time.Duration
	if outcomingRandom {
		timeToWait = (time.Second * time.Duration(rand.Intn(outcomingInterval)+1))
	} else {
		timeToWait = time.Second * time.Duration(outcomingInterval)
	}
	time.Sleep(timeToWait)
	b.available <- position
	b.resources[position] = false
	b.print()
}

func (b *Balancer) print() {
	//fmt.Print("\033[H\033[2J")
	for i := range b.resources {
		if b.resources[i] {
			color.Red.Print("●")
		} else {
			color.Green.Print("●")
		}
	}
	fmt.Println()
}

func (b *Balancer) run() {
	for n := 0; ; n++ {
		<-b.incoming
		var position int = <-b.available
		b.resources[position] = true
		b.print()
		go b.outcoming(n, position)
	}
}

func main() {
	var b Balancer = CreateBalancer(size)
	go b.runIncoming()
	b.run()
}
