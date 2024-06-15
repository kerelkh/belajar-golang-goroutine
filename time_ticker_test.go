package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

//ticker => representasi kejadian yang berulang
//jika duration 2 detik, maka channel akan dikirim setiap 2 detik
//membuat ticker timer.NewTicker(duration)
//untuk stop ticker Ticker.Stop()

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)

	stop := make(chan struct{})

	//untuk menghentikan ticker pake Stop()
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		close(stop)
	}()

	//.Stop() tidak menutup menunggu channel dari Tick sehingga function tetap berjalan terus
	// for tick := range ticker.C {
	// 	fmt.Println(tick)
	// }

	//ini jika menggunakan select
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println(tick)
		case <-stop:
			fmt.Println("stop ticker")
			return
		}
	}

}

// Tick untuk mendapatkan channelnya saja
func TestTick(t *testing.T) {
	counter := 0
	tick := time.Tick(2 * time.Second)
	for range tick {
		fmt.Println("Tick")
		counter++
		if counter > 5 {
			break
		}
	}

}
