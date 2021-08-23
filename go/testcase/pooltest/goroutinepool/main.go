package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Println("Hello world")
}

func demoFunc() {
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Hello world")
}

func main() {
	defer ants.Release()

	runTimes := 30
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("Running goroutines: %d\n", ants.Running())
	fmt.Println("finish all tasks")

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines :%d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d", sum)

}
