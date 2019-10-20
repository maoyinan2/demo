package main

import (
	_ "demo/mcfg"
	"sync"
	"time"

	// "demo/mlog"
	"fmt"
	// "log"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	// log.Println("abc")
	// for i := 0; i < 10; i++ {
	// 	mlog.Printf("test", "hello, mao%d", i)
	// }

	/* bytes, strings, strconv, unicode */
	var strs []string
	var str string = "Goodbye"

	fmt.Println(str + ",,")

	strs = append(strs, str) // _ = append 会失败
	fmt.Println(strs)

	tk := time.NewTicker(5 * time.Minute)
	go func(tk *time.Ticker) {
		defer wg.Done()

		<-tk.C
		fmt.Println("tick...")
	}(tk)

	wg.Wait()

}
