package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatal(err)
	}
	for{
		conn, err := ln.Accept()
		if err != nil{
			log.Fatal(err)
			conn.Close()
		}
	
		req, err  := bufio.NewReader(conn).ReadString('\n')
		if err != nil{
			log.Fatal(err)
		}
	
		fmt.Print(string(req))
	}
	
}
