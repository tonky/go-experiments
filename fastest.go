package main

import "fmt"

// import "math/rand"

import "time"

func main() {
	res := make(chan int)
	delays := []int{112, 59, 800, 44, 23, 500}

	fmt.Println(time.Now())

	for _, r := range delays {
		go State{r}.Work(res)
	}

	fmt.Println("\n", <-res)
}

type State struct {
	duration int
}

type Worker interface {
	Work(duration, res chan int)
}

func (s State) Work(res chan int) {
	fmt.Println("working for ", s.duration)
	time.Sleep(time.Duration(s.duration) * time.Millisecond)
	res <- s.duration
}
