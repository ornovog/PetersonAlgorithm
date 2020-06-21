package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func PetersonAlgorithmAtomic(n int) int{
	var wg sync.WaitGroup
	wg.Add(2)

	counter := 0
	var aWant, bWant,turn int32

	go func() {
		defer wg.Done()

		for i:=0; i<n/2; i++ {
			atomic.StoreInt32(&aWant,1)
			atomic.StoreInt32(&turn ,2)

			for ;atomic.LoadInt32(&bWant) == 1 && atomic.LoadInt32(&turn) == 2;{
			}
			counter++
			atomic.StoreInt32(&aWant,0)
		}
	}()

	go func() {
		defer wg.Done()

		for i:=0; i<n/2; i++ {
			atomic.StoreInt32(&bWant,1)
			atomic.StoreInt32(&turn ,1)
			for ;atomic.LoadInt32(&aWant) == 1 && atomic.LoadInt32(&turn) == 1;{
			}
			counter++
			atomic.StoreInt32(&bWant,0)
		}
	}()

	wg.Wait()
	return counter
}

func PetersonAlgorithm(n int)int{
	var aWants, bWants bool
	var turn string
	var counter int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i:=0;i<n/2;i++ {
			//A wants to enter the critical section
			aWants = true
			//A is setting turn to B
			turn = "B"
			//A is waiting for his turn
			for bWants && turn == "B" {
				time.Sleep(time.Nanosecond)
			}
			//Now this is A's turn in the critical section
			counter++
			aWants = false
			//A left the critical section
		}
	}()

	go func() {
		defer wg.Done()
		for i:=0;i<n-(n/2);i++ {
			//B wants to enter the critical section
			bWants = true
			//B is setting turn to A
			turn = "A"
			//B  is waiting for his turn
			for aWants && turn == "A" {
				time.Sleep(time.Nanosecond)
			}
			counter++
			bWants = false
			//B left the critical section
		}
	}()

	wg.Wait()

	return counter
}
