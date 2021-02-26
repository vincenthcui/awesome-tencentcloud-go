package main

import "sync"

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(1)
	if err := genaction(listenSignal()); err != nil {
		panic(err)
	}
	wg.Wait()
}
