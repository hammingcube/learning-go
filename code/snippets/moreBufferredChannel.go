package main

import "fmt"

func main() {
	c := make(chan int, 2)
	go func() {
		defer close(c)
		c <- 1
		c <- 2
		c <- 3
		c <- 4
	}()
	
	for i := range c {
		fmt.Println(i)
	}
}
