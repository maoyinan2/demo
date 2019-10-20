package main

import (
	"fmt"
	"sync"
	"time"
)

/* ticker 定期执行
timer只执行一次，time.Reset重置
time.Format("0102")，0101就不行
*/
func main() {
	// var a time.Time = time.Now()
	// s := a.Format("0101")
	// fmt.Println(s)

	// fmt.Println(time.Now().Format("0101"))

	var wg sync.WaitGroup
	wg.Add(2)

	timer1 := time.NewTimer(2 * time.Second)
	ticker1 := time.NewTicker(4 * time.Second)
	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer.")
		}
	}(timer1)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker.")
		}
	}(ticker1)

	wg.Wait()
}
