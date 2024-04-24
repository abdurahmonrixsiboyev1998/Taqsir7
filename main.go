package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	res := make(chan int)
	var wg sync.WaitGroup
	
	doTask := func(ID int) {
		defer wg.Done()
		time.Sleep(time.Duration(ID) * time.Second)
		res <- ID                                  
	}

	wg.Add(3)

	go doTask(1)
	go doTask(2)
	go doTask(3)

	go func() {
		wg.Wait() 
		close(res) 
	}()

	for res := range res {
		fmt.Println("Completed:", res)
	}
}
