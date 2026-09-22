package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Global variables shared between functions --A BAD IDEA
var i = 0

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, mut *sync.Mutex, Count int, sem *chan struct{}) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)  //all As should run before the Bs
	//Rendezvous here - implementing a barrier to account for 5 threads being called, don't think it would work as a rendezvous?
	mut.Lock() //gets unlocked after increment and check done
	i++
	if i == Count { //won't trigger until the last thread is coming through
		mut.Unlock()

		*sem <- struct{}{} //acts as release
	} else {
		mut.Unlock()
	}
	<-*sem //acts as wait
	i--
	if i != 0 { //will run for all but the last through so the last won't get stuck waiting, doubles as the lock for i--
		*sem <- struct{}{}
	}

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	sem := make(chan struct{}, 1)

	threadCount := 5 // should this be 2 for the rendezvous?
	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &mut, threadCount, &sem)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
