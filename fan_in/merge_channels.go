package fan_in

import (
	"sync"
)

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	kol_chan := len(channels)

	var wg sync.WaitGroup

	res_chan := make(chan int)

	wg.Add(kol_chan)

	go func() {
		wg.Wait()
		close(res_chan)
	}()

	for i := 0; i < kol_chan; i++ {
		go func(index int) {
			defer wg.Done()

			for v := range channels[index] {
				res_chan <- v
			}
		}(i)

	}

	return res_chan
}
