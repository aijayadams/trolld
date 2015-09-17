package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func LoadBundyBear() (bb []byte) {
	f, err := os.Open("bundy.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	filereader := bufio.NewReader(f)
	bb, err = filereader.ReadBytes('\x00')

	if err != io.EOF {
		log.Fatal(err)
		return
	}
	f.Close()
	return
}

func Telnet(c net.Conn, a Assets, conn_limit Counter) {
	defer c.Close()
	defer conn_limit.Remove()

	c.Write(a.trolls[0].asset)

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
