package main

var strChan = make(chan string, 3)

func main() {
	//----------------------demo7 桥接channel--------------------------
	// bridge := func(
	// 	done <-chan interface{},
	// 	chanStream <-chan <-chan interface{},
	// ) <-chan interface{} {
	// 	valStream := make(chan interface{})
	// 	go func() {
	// 		defer close(valStream)
	// 		for {
	// 			var stream <-chan interface{}
	// 			select {
	// 			case maybeStream, ok := <-chanStream:
	// 				if ok == false {
	// 					return
	// 				}
	// 				stream = maybeStream
	// 			case <-done:
	// 				return
	// 			}
	// 			// 通过桥接，我们可以在单个range 语句中使用处理channel 的channel ，并专注于我们的循环逻辑。
	// 			for val := range orDone(done, stream) {
	// 				select {
	// 				case valStream <- val:
	// 				case <-done:
	// 				}
	// 			}
	// 		}
	// 	}()
	// 	return valStream
	// }
	//----------------------demo7--------------------------

	//----------------------demo6 fanIn, fanOut--------------------------
	// 找10个素数，20多秒！
	// rand := func() interface{} { return rand.Intn(50000000) }

	// done := make(chan interface{})
	// defer close(done)

	// randIntStream := toInt(done, repeatFn(done, rand))
	// for prime := range take(done, primeFinder(done, randIntStream), 10) {
	// 	fmt.Printf("\t%d\n", prime)
	// }
	// =》 启动多个go程
	// numFinders := runtime.NumCPU()
	// finders := make([]<-chan int, numFinders)
	// for i := 0; i < numFinders; i++ {
	// 	finders[i] = primeFinder(done, randIntStream)
	// }
	// for prime := range take(done, fanIn(done, finders...), 10) {
	// }
	// =》 将结果汇总
	// fanIn := func(
	// 	done <-chan interface{},
	// 	channels ...<-chan interface{},
	// ) <-chan interface{} {
	// 	var wg sync.WaitGroup
	// 	multiplexedStream := make(chan interface{})
	// 	multiplex := func(c <-chan interface{}) {
	// 		defer wg.Done()
	// 		for i := range c {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case multiplexedStream <- i:
	// 			}
	// 		}
	// 	}

	// 	// 从所有channel取值
	// 	wg.Add(len(channels))
	// 	for _, c := range channels {
	// 		go multiplex(c)
	// 	}

	// 	// 等待所有读操作结束
	// 	go func() {
	// 		wg.Wait()
	// 		close(multiplexedStream)
	// 	}()

	// 	return multiplexedStream
	// }

	//----------------------demo6--------------------------

	//----------------------demo5 pipeline--------------------------
	// 流式处理或批处理数据。
	// generator := func(done <-chan interface{}, integers ...int) <-chan int {
	// 	intStream := make(chan int)
	// 	go func() {
	// 		defer close(intStream)
	// 		for _, i := range integers {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case intStream <- i:
	// 			}
	// 		}
	// 	}()
	// 	return intStream
	// }

	// multiply := func(
	// 	done <-chan interface{},
	// 	intStream <-chan int,
	// 	multiplier int,
	// ) <-chan int {
	// 	multipliedStream := make(chan int) // pineline的代价
	// 	go func() {
	// 		defer close(multipliedStream)
	// 		for i := range intStream {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case multipliedStream <- i * multiplier:
	// 			}
	// 		}
	// 	}()
	// 	return multipliedStream
	// }

	// add := func(
	// 	done <-chan interface{},
	// 	intStream <-chan int,
	// 	additive int,
	// ) <-chan int {
	// 	addedStream := make(chan int)
	// 	go func() {
	// 		defer close(addedStream)
	// 		for i := range intStream {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case addedStream <- i + additive:
	// 			}
	// 		}
	// 	}()
	// 	return addedStream
	// }

	// done := make(chan interface{})
	// defer close(done)

	// intStream := generator(done, 1, 2, 3, 4)
	// pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	// for v := range pipeline {
	// 	fmt.Println(v)
	// }
	//----------------------demo5--------------------------

	//----------------------demo4 or-channel--------------------------
	// 将一个或多个完成的channel 合并到一个完成的channel 中，该channel 在任何组件channel 关闭时关闭。
	// var or func(channels ...<-chan interface{}) <-chan interface{}
	// or = func(channels ...<-chan interface{}) <-chan interface{} {
	// 	switch len(channels) {
	// 	case 0:
	// 		return nil
	// 	case 1:
	// 		return channels[0]
	// 	}

	// 	orDone := make(chan interface{})
	// 	go func() {
	// 		defer close(orDone)

	// 		// 主体部分就是等待channel完成
	// 		switch len(channels) {
	// 		case 2:
	// 			select {
	// 			case <-channels[0]:
	// 			case <-channels[1]:
	// 			}
	// 		default:
	// 			select {
	// 			case <-channels[0]:
	// 			case <-channels[1]:
	// 			case <-channels[2]:
	// 			case <-or(append(channels[3:], orDone)...): // 递归等待
	// 			}
	// 		}
	// 	}()
	// 	return orDone
	// }

	// sig := func(after time.Duration) <-chan interface{} {
	// 	c := make(chan interface{})
	// 	go func() {
	// 		defer close(c)
	// 		time.Sleep(after)
	// 	}()
	// 	return c
	// }
	// start := time.Now()
	// <-or(
	// 	sig(2*time.Hour),
	// 	sig(5*time.Minute),
	// 	sig(time.Second),
	// 	sig(time.Hour),
	// 	sig(time.Minute),
	// )
	// fmt.Printf("done after %v", time.Since(start))
	//----------------------demo4--------------------------

	//----------------------demo3 go程泄露--------------------------
	// doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	// 	terminated := make(chan interface{})
	// 	go func() {
	// 		defer fmt.Println("doWork exited.")
	// 		defer close(terminated)
	// 		for {
	// 			select {
	// 			case s := <-strings:
	// 				fmt.Println(s)
	// 			case <-done:
	// 				return // 外层的 <-terminated 响应
	// 			}
	// 		}
	// 	}()
	// 	return terminated
	// }

	// done := make(chan interface{})
	// terminated := doWork(done, nil)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Canceling doWork goroutine...")
	// 	close(done) // 外层控制go程的退出
	// }()

	// <-terminated // 等待doWork完成
	// fmt.Println("Done")
	//----------------------demo3--------------------------

	//----------------------demo2 for-select模型--------------------------
	// intChan := make(chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	intChan <- i
	// }
	// close(intChan) // 没有这句会死锁！
	// syncChan := make(chan struct{}, 1)
	// go func() {
	// Loop:
	// 	for {
	// 		select {
	// 		case e, ok := <-intChan:
	// 			if !ok {
	// 				fmt.Println("end.")
	// 				break Loop // 跳过多层 select & for
	// 			}
	// 			fmt.Printf("Received: %v\n", e)
	// 		}
	// 	}
	// 	syncChan <- struct{}{}
	// }()
	// <-syncChan //阻塞main
	//----------------------demo2--------------------------

	//----------------------demo1 go程流程绘制--------------------------
	// syncChan1 := make(chan struct{}, 1)
	// syncChan2 := make(chan struct{}, 2)

	// go func() {
	// 	<-syncChan1
	// 	fmt.Println("Received a sync signal and wait a second...[receiver]")
	// 	time.Sleep(time.Second)
	// 	for {
	// 		if elem, ok := <-strChan; ok {
	// 			fmt.Println("Received:", elem, "[receiver]")
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println("Stopped.[receiver]")
	// 	syncChan1 <- struct{}{}
	// }()

	// go func() {
	// 	for _, elem := range []string{"a", "b", "c", "d"} {
	// 		strChan <- elem
	// 		fmt.Println("Send:", elem, "[sender]")
	// 		if elem == "c" {
	// 			syncChan1 <- struct{}{}
	// 			fmt.Println("Sent a sync signal.[sender]")
	// 		}
	// 	}
	// 	fmt.Println("Wait 2 seconds...[sender]")
	// 	time.Sleep(2 * time.Second)
	// 	close(strChan)
	// 	syncChan2 <- struct{}{}
	// }()
	// <-syncChan2
	// <-syncChan1
	//----------------------demo1--------------------------
}
