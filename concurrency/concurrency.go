package main

// In this example, the closure does a non-blocking send, which it achieves by using the send
// operation in select statement with a default case. If the send cannot go through
// immediately the default case will be selected. Making the send non-blocking guarantees
// that none of the goroutines launched in the loop will hang around. However, if the result
// arrives before the main function has made it to the receive, the send could fail since no
// one is ready.

// This problem is a textbook example of what is known as a race condition, but the fix is
// trivial. We just make sure to buffer the channel ch (by adding the buffer length as the
// second argument to make), guaranteeing that the first send has a place to put the
// value. This ensures the send will always succeed, and the first value to arrive will be
// retrieved regardless of the order of execution.

import "fmt"
import "time"

type Result int
type Conn int

func (c *Conn) DoQuery(params string) Result {
	time.Sleep(time.Second)
	return 1
}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result, 1)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
			}
		}(conn)
	}
	return <-ch
}

func main() {
	conns := []Conn{&Conn{}, &Conn{}}

	cq := &Conn{}
	cq.DoQuery("a")
}
