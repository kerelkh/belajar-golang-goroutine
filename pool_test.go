package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// POOL sudah mendukung kondisi race condition, sehingga tidak akan terkendala race condition
func TestPool(t *testing.T) {
	//New untuk default jika pool tidak ada data else nilai akan Nil
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	newData := []string{"kerel", "khalif", "afif"}
	for _, data := range newData {
		pool.Put(data)
	}

	for i := 1; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Pool ", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Proses selesai")
}
