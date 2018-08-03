package main

import (
	s "./source"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Launch Server...")

	listen, err := net.Listen("tcp", s.IP)
	s.CheckErr(err)

	for {
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println("Message:", err)
			continue
		}
		go handleServerConnection(connect)
	}
}

func handleServerConnection(connect net.Conn) {
	fmt.Println("Accepted connection")

	for {
		var msg int
		answer := s.Answer{}

		decoder := json.NewDecoder(connect)
		decod := decoder.Decode(&msg)
		s.CheckErr(decod)
		fmt.Println("Message for: ", msg)

		t := time.Now()
		answer.Answer = s.FibonacciResult(msg)
		answer.Time = time.Since(t)

		encoder := json.NewEncoder(connect)
		encod := encoder.Encode(answer)
		s.CheckErr(encod)
	}
	connect.Close()
}
