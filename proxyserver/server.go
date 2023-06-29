package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is up on 127.0.0.1:8080 ") // TODO: rework with argument
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			conn.Close()
		}

		req, err := bufio.NewReader(conn).ReadString('\n')
		req = strings.Replace(req, "\n", "", -1)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(req)
		// u, err := url.Parse(req)
		// if err != nil {
		// 	log.Fatal(err)
		// }
			
		resp, err := http.Get(req)
		if err != nil {
			log.Fatal(err)
		}
		
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Fprintf(conn, string(body))
		conn.Close()

	}

}
