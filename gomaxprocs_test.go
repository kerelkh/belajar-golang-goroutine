package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

//gomacprocs sebuah function package runtime utuk mengubah jumlah thread atau mengambil jumlah thread
//secara default, jumlah thread di Go-Lang itu sebangak jumlah CPU di komputer
//kita bisa melihat jumlah CPU dengan function runtime.NumCpu()

func TestGoMaxProcs(t *testing.T) {
	//mengambil jumlah CPU di komputer
	cpu := runtime.NumCPU()
	fmt.Println(cpu)

	//-1 untuk melihat total
	//diatas 0 untuk menambah
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	runtime.GOMAXPROCS(200)
	fmt.Println("total thread setelah ditambah", runtime.GOMAXPROCS(-1))

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)

	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			defer group.Done()
		}()
	}

	totalGoroutine = runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)
}
