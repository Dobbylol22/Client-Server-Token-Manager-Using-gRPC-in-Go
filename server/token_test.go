package main

import (
	"strconv"
	"sync"
	"testing"
)

func TestConcurrency(t *testing.T) {
	wg := &sync.WaitGroup{} //it waits for a collections of goroutines to finish.
	m := NewTokenManager()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go RunACycle(m, wg, strconv.Itoa(i), strconv.Itoa(i), uint64(i*10), uint64(i*20), uint64(i*30))
	}
	wg.Wait()
}

func RunACycle(m *TokenManager, wg *sync.WaitGroup, id string, name string, low uint64, mid uint64, high uint64) (err error) {
	defer wg.Done()
	err = m.Create(id)
	if err != nil {
		return
	}
	_, err = m.Write(id, name, low, mid, high)
	if err != nil {
		return
	}
	_, err = m.Read(id)
	if err != nil {
		return
	}
	err = m.Drop(id)
	if err != nil {
		return
	}
	return
}
