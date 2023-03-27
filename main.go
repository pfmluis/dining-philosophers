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
	rightFork *Fork
	leftFork  *Fork
}

func (p *Philosopher) Eat() {
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("%s is eating\n", p.name)
	time.Sleep((time.Duration(rand.Intn(3) + 1)) * time.Second)

	p.leftFork.Unlock()
	p.rightFork.Unlock()
	fmt.Printf("%s is back to thinking\n", p.name)
}

func main() {
	var wg sync.WaitGroup
	philosopherNames := []string{"Confucius", "Socrates", "Plato", "Descartes", "Kant"}
	forks := []*Fork{new(Fork), new(Fork), new(Fork), new(Fork), new(Fork)}
	for i, name := range philosopherNames {
		philosopher := Philosopher{
			name:      name,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%len(forks)],
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			philosopher.Eat()
		}()
	}

	wg.Wait()
	fmt.Println("Everyone is done eating!")
}
