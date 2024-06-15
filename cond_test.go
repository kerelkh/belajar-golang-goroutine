package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

//Cond => implementasi locking dengan kondisi

func WaitCondition(value int) {
	defer group.Done()

	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}
func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	//tanpa signal akan terkena deadlock
	//dibawal akan memberi signal tiap 1 detik sehingga goroutine dijalankan 1 per 1 detik
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()

		}
	}()

	//broadcast, setelah check wait maka semuanya diminta langsung jalan
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()

}
