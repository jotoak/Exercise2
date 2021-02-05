// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)



func incrementer(acess chan int, finished chan<- bool, i *int) {
	
	for j := 0; j < 1000000; j++ {
		turn :=<-acess
		(*i)--
		acess <-turn
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementer(acess chan int, finished chan<- bool,i *int) {
	
	for j := 0; j < 1000000+1; j++ {
		turn :=<-acess
		(*i)--
		acess <-turn
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the remaining channels
	acess := make(chan int,1)
	finished := make(chan bool)
	i :=0
	acess <-1
	go incrementer(acess,finished,&i)
	go decrementer(acess,finished, &i)
	// TODO: Spawn the required goroutines

	// TODO: block on finished from both "worker" goroutines

	fmt.Println("The magic number is:", i)
}
