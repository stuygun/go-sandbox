package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

var c = make(chan Philo, 2)

type ChopS struct{ sync.Mutex }

type Philo struct {
	times           int
	number          int
	leftCS, rightCS *ChopS
}

func getPermission(philo Philo) bool {
	if philo.times < 2 {
		philo.times += 1
		c <- philo
		return true
	}
	return false
}

func (p Philo) eat() {
	if getPermission(p) {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat  " + strconv.Itoa(p.number))
		time.Sleep(time.Second)
		fmt.Println("finishing eating  " + strconv.Itoa(p.number))

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-c
		waitGroup.Done()
	}

}
func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{0, i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	waitGroup.Add(15)
	for j := 0; j < 3; j++ {

		for i := 0; i < 5; i++ {
			go philos[i].eat()
		}
	}
	waitGroup.Wait()
}
