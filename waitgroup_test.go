package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// waitGroup dipakai untuk menunggu goroutine selesai dijalankan yang seblumnya kita pakai seperti time.Sleep
// waitgroup seperti counter jika Add() akan ditambah, jika Done() maka dikurang,
// jika counter 0 maka program tidak akan selesai (masih menunggu)
func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello World")
	time.Sleep(5 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Goroutine selesai dijalankan")
}
