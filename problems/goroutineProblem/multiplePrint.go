package goroutineproblem

import (
	"fmt"
	"strconv"
	"sync"
)

/*
   Print 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728 via goroutine
*/

// wg := sync.WaitGroup{}
var wg sync.WaitGroup

func printChar(c chan bool) {
	// 65~90 : upper char
	i := 65
	for {
		select {
		case v := <-c:

			if !v {
				wg.Done()
				fmt.Print("\n")
				return
			}
			r := string(rune(i)) + string(rune(i+1))
			fmt.Print(r)
			i = i + 2

			c <- true
		}
	}
}

func printNum(c chan bool) {
	fmt.Print("12")
	c <- true
	s := 3
	for {
		select {
		case <-c:
			r := strconv.Itoa(s) + strconv.Itoa(s+1)
			fmt.Print(r)
			s = s + 2
			if s > 28 {
				wg.Done()
				c <- false
			}
			c <- true
		}
	}

}

/* Usage
func main() {
	c := make(chan bool)
	wg.Add(2)
	go printChar(c)
	go printNum(c)
	wg.Wait()
	runtime.Gosched()

}*/
