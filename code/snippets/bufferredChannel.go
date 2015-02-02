package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	close(c)
	
	for i := range c {
		fmt.Println(i)
	}
}
