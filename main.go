package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	quit := make(chan struct{})
	fnum := make(chan int)

	go printFibonacci(done, fnum)

	go fiboGuess(done, quit, fnum)

	<-quit
}

func fiboGuess(done, quit chan struct{}, fnum chan int) {
	var n, q int
	go func() {
		for q = range fnum {
		}
	}()
	for {
		fmt.Scanln(&n)
		if isFibonacci(n, q) {
			fmt.Println("Cool! You are better than this program.")

			done <- struct{}{}
			go spamThanks(quit)

			return
		}
	}
}

func isFibonacci(n, q int) bool {
	return n == fibonacci(q+1)
}

func printFibonacci(done chan struct{}, fnum chan int) {
	for i := 1; ; i++ {
		time.Sleep(time.Second * 1)
		select {
		case <-done:
			close(fnum)
			return
		default:
			res := fibonacci(i)
			fnum <- i
			fmt.Println("Fibonacci number ", i, "=", res)
		}
	}
}

func fibonacci(n int) int {
	first, second := 0, 1
	for i := 1; i < n; i++ {
		first, second = second, first+second
	}
	return first
}

func spamThanks(quit chan struct{}) {
	var stopWord string

	go func() {
		for {
			time.Sleep(time.Second * 1)
			select {
			case <-quit:
				return
			default:
				fmt.Println("Молодец")
			}
		}
	}()

	for {
		fmt.Scanln(&stopWord)
		if stopWord == "thx" || stopWord == "спс" {
			fmt.Println("Спасибо, до свидания!")
			quit <- struct{}{}
			return
		}
	}
}
