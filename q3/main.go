package main

import("fmt"
"time")

func goroutine(value chan int){
	fmt.Println("Goroutine begins")
	value <- 5
	time.Sleep(1*time.Second)
	fmt.Println("Goroutine finish")
	close (value)
}

func main(){
	value := make(chan int)
	fmt.Println("Main executes")
	go goroutine(value)
	fmt.Println("Main routine")
	time.Sleep(3 * time.Second)
	result := <- value
	fmt.Println(result)
	fmt.Println("Done")
}
