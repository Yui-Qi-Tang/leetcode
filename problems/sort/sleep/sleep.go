package sleep

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSleep(v uint64) {
	time.Sleep(time.Duration(v) * time.Second)
}

// Sort sorting here, consider positve integer firstly; statify stable property in this practice
func Sort(input []uint64) (result []uint64) {
	for _, v := range input {
		wg.Add(1)
		go func(v uint64) {
			defer wg.Done()
			doSleep(v)
			result = append(result, v)
		}(v)
	}
	wg.Wait()
	return
}
