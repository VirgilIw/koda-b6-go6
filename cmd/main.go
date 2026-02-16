package main

import (
	"sync"

	"github.com/virgilIw/koda-b6-go6/internal"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan internal.Data)
	wg.Add(2)
	go internal.Queque(ch, &wg)
	go internal.ProcessQueque(ch, &wg, &mu)
	//
	wg.Wait()
}
