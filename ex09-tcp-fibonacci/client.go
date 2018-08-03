package main

import (
	s "./source"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {
	connect, err := net.Dial("tcp", s.IP)
	s.CheckErr(err)

	defer connect.Close()

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input := s.Reader(scan.Text())

		encoder := json.NewEncoder(connect)
		encod := encoder.Encode(input)
		s.CheckErr(encod)

		var msg s.Answer
		decoder := json.NewDecoder(connect)
		decod := decoder.Decode(&msg)
		s.CheckErr(decod)

		fmt.Printf("%s %d\n", msg.Time, msg.Answer)
	}
}
