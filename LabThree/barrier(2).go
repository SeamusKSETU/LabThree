//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Seamus Kennedy C00209242
//collaborated with Sam C and Bartosz
// Issues:
// The barrier is not implemented!
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var j = 0
var ctx = context.Background() //moved to save passing pointers
var sem = semaphore.NewWeighted(0)
var barrier = make(chan struct{}, 1)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, theLock *sync.Mutex, tot int) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	theLock.Lock()
	j++
	if j == tot {
		fmt.Println(goNum, "asd") //troubleshooting
		theLock.Unlock()
		//sem.Release(1) for some reason this was crashing it as soon as called, might have misunderstood how they work, using channel instead
		barrier <- struct{}{}
		//fmt.Println("rel")
	} else {
		fmt.Println(goNum, " tuh", j)
		theLock.Unlock()
	}
	//sem.Acquire(ctx, 1)
	<-barrier
	//fmt.Println(goNum, " poi")
	j--
	if j != 0 {
		//fmt.Println(goNum, " tyu", j)
		//sem.Release(1)
		barrier <- struct{}{}
	}

	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these

	var theLock sync.Mutex
	theLock.Lock()
	j = 0
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &theLock, totalRoutines)
	}
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
