package channels

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) { //, done chan bool) {
	fmt.Println("❯ worker [LB]", workerId)
	for res := range msg {
		fmt.Println("Worker [LB]", workerId, "Msg", res)
		time.Sleep(time.Second)
		// done <- true // signal that the task is done
	}
}

func loadBalancer() chan int {
	fmt.Println("❯ channel load balancer")

	msgChannel := make(chan int)
	qtdWorkers := 7
	// done := make(chan bool) // channel for signaling when all workers are done

	// create workers
	for i := range qtdWorkers {
		go worker(i, msgChannel) // , done)
	}

	// send messages to workers
	for i := 0; i < 10; i++ {
		msgChannel <- i
	}

	return msgChannel
}
