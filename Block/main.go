package main

import (
	"fmt"
	"time"
)

func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
	}

	close(c)
}

func receive(c <-chan int) {
	for {
		a, ok := <-c
		if ok == false {
			fmt.Println("끝.")
			break
		}

		fmt.Printf("뭘 받았나? : %d\n", a)
	}
}

func main() {
	c := make(chan int)
	go countToTen(c)
	receive(c)

}
