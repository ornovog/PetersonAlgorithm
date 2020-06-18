package main

import (
	"testing"
	"time"
)

var N = 10000000

func TestPetersonAlgorithm(t *testing.T) {
	c1 := make(chan int, 1)

	go func() {
		c1 <- PetersonAlgorithm(N)
	}()

	select {
	case result := <-c1:
		if result != N  {
			t.Errorf("result equals to %v",result)
		}
	case <-time.After(3 * time.Second):
			t.Errorf("got timeout")
	}
}

func TestPetersonAlgorithmAtomic(t *testing.T) {
	c1 := make(chan int, 1)

	go func() {
		c1 <- PetersonAlgorithmAtomic(N)
	}()

	select {
	case result := <-c1:
		if result != N  {
			t.Errorf("result equals to %v",result)
		}
	case <-time.After(3 * time.Second):
		t.Errorf("got timeout")
	}
}
