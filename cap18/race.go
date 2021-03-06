package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	contador := 0
	wgCount := 10

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutine (1)", runtime.NumGoroutine())

	wg.Add(wgCount)

	for i := 0; i < wgCount; i++ {
		go func() {
			v := contador

			//yield
			runtime.Gosched()

			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("GoRoutine (2)", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println(contador)
	fmt.Println("GoRoutine (3)", runtime.NumGoroutine())
}
