package main

import (
	"fmt"
	"time"
)

func main() {
	alice := Person{15, "Alice"}
	fmt.Println(alice.name) // Call String method of p, of type pair.

	check := alice.is_underage()

	fmt.Printf("Person: %s, is_underage: %s \n", alice.name, check)

	c := make(chan bool)

	go func(ch chan bool) {
		time.Sleep(1 * time.Second)
		ch <- false
	}(c)

	select {
	case val := <-c:
		fmt.Printf("got: ", val)
	case <-time.After(1200 * time.Millisecond):
		fmt.Printf("timeout")
	}

}

/*
func is_underage(p Person) bool {
	if p.age < 18 {
		return true
	}

	return false
}
*/
type Person struct {
	age  int
	name string
}

func (p Person) is_underage() bool {
	if p.age < 18 {
		return true
	}

	return false
}

/*
type Person interface {
	func is_underage(p Person) bool {
		if p.age < 18 {
			return true
		}

		return false
	}
}
*/
