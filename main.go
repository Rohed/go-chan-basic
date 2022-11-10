package main

import (
	"fmt"
	"time"
)

 
func main() {
	ch := make(chan int,2)
	go hold(ch)

	time.Sleep(2 * time.Second)
	fmt.Println(time.Now(), "Waiting for messages")

	fmt.Println(time.Now(), "received 1", <-ch)
	fmt.Println(time.Now(), "received 2", <-ch)
	fmt.Println(time.Now(), "received 3", <-ch)

	fmt.Println(time.Now(), "Exiting")
}


func hold(ch chan int){
	for i:= 0; i< 3; i++{
		fmt.Println(time.Now(),i,"sending");
		ch <- i
		fmt.Println(time.Now(),i,"sent");
	}
}