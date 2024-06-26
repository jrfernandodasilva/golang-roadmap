package channels

import (
	"fmt"
	"runtime"
	"time"
)

func couter(mode string) {
	for i := 0; i < 5; i++ {
		fmt.Println(mode, i)
		time.Sleep(time.Second)
	}
}

func singleProcess() {
	runtime.GOMAXPROCS(1)
	fmt.Println("❯ single process - start")

	go func() {
		// go skips loop - predictive optimization
		for {
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("❯ single process - end")
}

func channel() {
	// channel is a synchronization primitive
	hello := make(chan string)

	// goroutine can send to channel
	go func() {
		hello <- "Hello World"
	}()

	// goroutine can receive from channel
	msg := <-hello
	fmt.Println("❯ channel")
	fmt.Println(msg)
}

func channelSelect() {
	hello := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		hello <- "Hello World"
	}()

	select {
	case msg := <-hello:
		fmt.Println("❯ channel select")
		fmt.Println(msg)
	default:
		fmt.Println("❯ channel select - default")
	}

	time.Sleep(time.Second * 5)
}

func channelSenderReceiver() {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			queue <- i
			i++
		}
		// close(queue)
	}()

	fmt.Println("❯ channel sender - receiver")
	for x := range queue {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}

func channelForever() {
	forever := make(chan string)

	go func() {
		x := true
		for {
			if x {
				continue
			}
		}
	}()

	fmt.Println("❯ waiting channel forever")
	<-forever
}

func workerForever(workerId int, msg chan int) {
	fmt.Println("❯ workerForever", workerId)
	for { // infinite loop
		res, ok := <-msg // receive messa from channel
		if !ok {         // check if channel is closed
			break // exit loop if channel is closed
		}
		fmt.Println("workerForever", workerId, "Msg", res)
		time.Sleep(time.Second)
	}
}

func Run() {
	fmt.Println("=====================================")
	fmt.Println("===============channels==============")
	fmt.Println("=====================================")

	fmt.Println("❯ channels start")

	msgChannel := loadBalancer()
	bufferedChannel()
	channelSynchronization()

	go couter("goroutine-a")
	go couter("goroutine-b")

	time.Sleep(time.Second * 5)

	singleProcess()
	channel()
	channelSelect()

	msgForever := make(chan int)
	go workerForever(44, msgForever) // , done)
	go workerForever(55, msgForever) // , done)
	for i := 0; i < 6; i++ {
		msgForever <- i
	}

	// for i := 1; i <= 7; i++ {
	// 	<-done
	// }
	// close(mgs)

	go func() {
		time.Sleep(time.Second * 2)
		msgForever <- 99
		msgChannel <- 15
		time.Sleep(time.Second * 5)
		msgForever <- 77
		close(msgChannel)
		close(msgForever)
	}()

	// channelSenderReceiver()
	// channelForever()

	fmt.Println("❯ channels end")
}
