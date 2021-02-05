
package main

import "fmt"
import "time"


func producer(link chan int,done chan<- bool ){

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        link <- i
    }
    done <-true
}

func consumer(link chan int){

    time.Sleep(1 * time.Second)
    for {
        i := <-link
         //TODO: get real value from buffer
        fmt.Printf("[consumer]: %d\n", i)
        time.Sleep(50 * time.Millisecond)
    }

}


func main(){
    
    // TODO: make a bounded buffer
    link:=make(chan int)
    done:=make(chan bool)
    go consumer(link)
    go producer(link,done)
    select {}
}