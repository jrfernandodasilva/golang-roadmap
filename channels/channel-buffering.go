package channels

import "fmt"

// buffered channel is a channel with a buffer
func bufferedChannel() {

	fmt.Println("❯ buffered channel")
	msg := make(chan string, 2)

	msg <- "[0]buffered"
	msg <- "[1]channel"
	close(msg)

	fmt.Println(<-msg)
	fmt.Println(<-msg)

	// channel closed - no more messages
	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
