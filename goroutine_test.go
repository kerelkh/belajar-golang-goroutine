package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

func TestGoroutineHelloWorld(t *testing.T) {
	//goroutine (proses asynchronus)
	go RunHelloWorld()
	//--
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(i int) {
	fmt.Println(i)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}

func channelData(ch chan int, i int) {
	ch <- i
	defer close(ch)
}
func TestChannel(t *testing.T) {
	ch := make(chan int)
	go channelData(ch, 10)

	data := <-ch

	fmt.Println(data)
}

//channel in-out, memberitahu bahwa function digunakan untuk mengirim atau menerima data

func OnlyIn(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello World!"
}

func OnlyOut(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func TestChannelOnlyInOut(t *testing.T) {
	ch := make(chan string)
	go OnlyIn(ch)
	go OnlyOut(ch)
	time.Sleep(3 * time.Second)
}

// buffered channel => memungkinkan untuk menampung data jika penerima masih belum tersedia
func TestBufferedChannel(t *testing.T) {

	//kapasitas antrian 5 untuk channel
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	fmt.Println(cap(ch))
	fmt.Println(len(ch))

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

		time.Sleep(4 * time.Second)
	}()
}

func TestBufferedDeadlock(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2 // deadlock karena channel sudah penuh

	close(ch)
	fmt.Println(<-ch)

	fmt.Println("Selesai")
}

// range channel jika tidak tau berapa banyak yang dikirimkan oleh pengirim data pada channel
func TestRangeChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}
}

// select channel memungkinkan mengambil dari beberapa channel yang paling cepat mengirimkan data
func GiveMeResponse(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello"
	defer close(ch)
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	for counter := 0; ; {
		select {
		case data := <-ch1:
			fmt.Println("Data dari ch1", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari ch2", data)
			counter++
		}

		if counter >= 2 {
			break
		}
	}
}

// default channel jika belum menerima data dari channel
func TestSelectDefaultChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	for counter := 0; ; {
		select {
		case data := <-ch1:
			fmt.Println("Data dari ch1", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari ch2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter >= 2 {
			break
		}
	}
}
