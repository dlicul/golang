package main

import (
	"strconv"
)

func send(messages chan<- string, msg string) {
	messages <- "ping" + msg
}

func printMsgs(messages <-chan string, finished chan<- bool) {
	for {
		msg, ok := <-messages
		if ok {
			println("message received=", msg)
		} else {
			println("no more messages")
			finished <- true
			return
		}
	}
}

func main() {
	var messages = make(chan string)
	var finished = make(chan bool)

	go printMsgs(messages, finished)
	for i := 1; i <= 10; i++ {
		send(messages, strconv.Itoa(i))
		//println("sent", i)
	}

	close(messages)
	println("---sent all---")

	<-finished // wait for messages to be received
}
