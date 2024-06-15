package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// bahayanya concurency karna akses sesuatu secara bersamaan, sehingga menimbulkan masalah race condition
func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}

// untuk mengatasi masalah bisa gunakan sync.Mutex untuk proses Lock dan Unlock
func TestMutexRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}
