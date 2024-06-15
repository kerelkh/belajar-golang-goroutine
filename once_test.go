package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func OnlyOnce(counter *int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	*counter++
}

// ONCE digunakan untuk menjalankan fungsi hanya sekali, sehingga fungsi di goroutine lain akan diabaikan
func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	counter := 0

	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				OnlyOnce(&counter, group)
			})
		}()
	}

	group.Wait()
	fmt.Println("Goroutine berhasil dijalankan dengan counter:", counter)
}
