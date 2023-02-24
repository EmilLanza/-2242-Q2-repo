package main

import("fmt"
"time")

func goroutine(){
	fmt.Println("Goroutine begins")
	time.Sleep(1*time.Second)
	fmt.Println("Goroutine finish")
}

func main(){
	fmt.Println("Main executes")
	go goroutine()
	fmt.Println("Main routine")
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}