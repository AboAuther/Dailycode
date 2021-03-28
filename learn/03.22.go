package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d", i)
	}
}
func main() {
	runtime.GOMAXPROCS(5)
	go a()
	go b()
	time.Sleep(time.Second)
}
