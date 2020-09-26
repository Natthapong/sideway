package main

import (
	"fmt"
	"sync"
)

func say(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello,", name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go say("natthapong", &wg)

	fmt.Println("before wait !!!")
	wg.Wait()
	fmt.Println("after wait !!!")

}
