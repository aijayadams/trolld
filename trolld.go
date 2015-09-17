package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// Counter for locking and limits
type Counter struct {
	sync.Mutex

	count int
	Max   int
}

// Constructor for locking counters
func NewCounter(max int) Counter {
	return Counter{Max: max}
}

// Increment locking counters
func (c Counter) Add() error {
	c.Lock()
	defer c.Unlock()
	if c.count < c.Max {
		c.count++
		return nil
	}
	return fmt.Errorf("Connection Limit Exceeded")
}

// Decrement locking counters
func (c Counter) Remove() {
	c.Lock()
	defer c.Unlock()
	c.count = c.count - 1
}

func main() {
	// Set connection limit
	conn_limit := NewCounter(16)

	//
	bear := LoadBundyBear()
	l, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		// Check Connection Limit
		if e := conn_limit.Add(); e != nil {
			time.Sleep(1)
		}
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go Telnet(conn, bear, conn_limit)
	}
}
