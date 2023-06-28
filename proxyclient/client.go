package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		conn.Close()
	}

	read := bufio.NewReader(os.Stdin)
	st, err := read.ReadString('\n')
	fmt.Fprintf(conn, st)
	
}
