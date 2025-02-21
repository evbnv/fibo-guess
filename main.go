package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	quit := make(chan struct{})

	go printFibonacci(done)

	go fiboGuess(done, quit)

	<-quit
}

func fiboGuess(done, quess chan struct{}) {
	var n int
	for {
		fmt.Scan(&n)
		if isFibonacci(n) {
			fmt.Println("Congratulations! You guessed the correct Fibonacci number.")
			done <- struct{}{}

			go spamThanks(quess)

			return
		} else {
			fmt.Println("Sorry, that's not a Fibonacci number. Try again.")
		}
	}
}

func printFibonacci(done chan struct{}) {
	for i := 1; ; i++ {
		select {
		case <-done:
			return
		default:
			time.Sleep(time.Second * 3)
			fmt.Println("Fibonacci number ", i, "=", fibonacci(i))
		}
	}
}

func fibonacci(n int) int {
	first, second := 1, 1
	for i := 1; i < n; i++ {
		first, second = second, first+second
	}
	return first
}

func isFibonacci(n int) bool {
	for i := 1; ; i++ {
		fib := fibonacci(i)
		if fib > n {
			return false
		}
		if fib == n {
			return true
		}
	}
}

func spamThanks(quit chan struct{}) {
	var stopWord string

	go func() {
		select {
		case <-quit:
			return
		default:
			for {
				fmt.Scan(&stopWord)
			}
		}
	}()

	for stopWord != "thx" && stopWord != "спс" {
		time.Sleep(time.Second * 2)
		fmt.Println("Молодец")
	}
	fmt.Println("Спасибо, до свидания!")
	quit <- struct{}{}
}
