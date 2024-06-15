package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// timer 5 detik
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

// after func digunakan untuk menjalankan delay function (delay job)
func TestAfterFunc(t *testing.T) {
	group := &sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Executed after 2 second")
		group.Done()
	})

	group.Wait()
	fmt.Println("Proses selesai")
}

func TestAfterFunc2(t *testing.T) {
	group := &sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Executed after 5 second at ", time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
