package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//atomic digunakan untuk menggunakan data primitive secara aman pada proses concurrent
//sehingga Mutex tidak diperlukan

func TestAtomic(t *testing.T) {
	counter := int64(0)
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter: ", counter)
}
