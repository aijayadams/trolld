package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
	"math/rand"

)

func Telnet(c net.Conn, a Assets, conn_limit Counter) {
	defer c.Close()
	defer conn_limit.Remove()

	// Select troll and deploy
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	troll_int := r.Int31n(int32(a.Len()))

	fmt.Println("Troll: " + a.trolls[troll_int].name)
	c.Write(a.trolls[troll_int].asset)

	buf := bufio.NewReader(c)
	for {
		_, _, err := buf.ReadLine()
		if err == io.EOF {
			fmt.Println("Connection Closed")
			return
		}
		c.Close()
	}
}
