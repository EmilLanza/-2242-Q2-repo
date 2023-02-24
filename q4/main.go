package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func count(){
	for i := 0; i < 10; i++{
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}
	wg.Done()
}

func main(){
	wg.Add(1)
	go count()
	wg.Wait()
	
}