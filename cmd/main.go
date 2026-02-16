package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/virgilIw/koda-b6-go6/internal"
)

// mutex itu mengunci akses ke bagian kode tertentu supaya hanya satu goroutine yang bisa masuk pada saat yang sama.

func processQueue(ch chan internal.Data, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for data := range ch {

		mu.Lock()
		fmt.Printf("Yang Memesan: %s\n", data.Name)
		mu.Unlock()

		for i := data.Wait; i > 0; i-- {
			mu.Lock()
			fmt.Printf("Tunggu %d detik...\n", i)
			mu.Unlock()

			time.Sleep(1 * time.Second)
		}

		mu.Lock()
		fmt.Println("Selesai:", data.Name)
		fmt.Println("------------------")
		mu.Unlock()
	}
}

func main() {
	ch := make(chan internal.Data)

	var wg sync.WaitGroup
	var mu sync.Mutex

	go internal.Queue(ch)

	wg.Add(1)
	go processQueue(ch, &wg, &mu)

	wg.Wait()
}
