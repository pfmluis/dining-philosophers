package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fork struct{ sync.Mutex }

type Philosopher struct {
	name      string
	waitGroup *sync.WaitGroup
	rightFork *Fork
	leftFork  *Fork
}

func (p *Philosopher) Eat() {
	defer p.waitGroup.Done()
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("%s is eating\n", p.name)
	time.Sleep((time.Duration(rand.Intn(3) + 1)) * time.Second)

	p.leftFork.Unlock()
	p.rightFork.Unlock()
	fmt.Printf("%s is back to thinking\n", p.name)
}

func main() {
	var waitGroup sync.WaitGroup
	philosopherNames := []string{"Confucius", "Socrates", "Plato", "Descartes", "Kant"}
	forks := make([]*Fork, len(philosopherNames))

	for i := range forks {
		forks[i] = new(Fork)
	}

	for i, name := range philosopherNames {
		philosopher := Philosopher{
			name:      name,
			waitGroup: &waitGroup,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%len(forks)],
		}

		waitGroup.Add(1)
		go philosopher.Eat()
	}

	waitGroup.Wait()
	fmt.Println("Everyone is done eating!")
}
