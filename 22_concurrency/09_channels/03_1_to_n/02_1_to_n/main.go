package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			num := i
			for n := range c {
				fmt.Println("goroutine number:", num, " value:", n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
