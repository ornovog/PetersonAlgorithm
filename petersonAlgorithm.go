package main

import (
	"sync"
	"sync/atomic"
)


func PetersonAlgorithm(n int) int{
	var wg sync.WaitGroup
	wg.Add(2)

	counter := 0
	var aWant, bWant bool
	var turn string

	go func() {
		defer wg.Done()

		for i:=0; i<n/2; i++ {
			aWant = true
			turn ="b"

			for ;bWant && turn=="b";{
			}
			counter++
			aWant=false
		}
	}()

	go func() {
		defer wg.Done()

		for i:=0; i<n/2; i++ {
			bWant = true
			turn ="a"

			for ;aWant && turn=="a";{
			}
			counter++
			bWant=false
		}
	}()

	wg.Wait()
	return counter
}

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