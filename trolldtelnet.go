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

func Telnet(c net.Conn, bear []byte, conn_limit Counter) {
	fmt.Println(c.RemoteAddr())
	defer c.Close()
	defer conn_limit.Remove()

	c.Write(bear)

	buf := bufio.NewReader(c)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			fmt.Println("esmf")
			return
		}
		fmt.Println(string(line))
		c.Write(line)
		c.Close()
	}
}
