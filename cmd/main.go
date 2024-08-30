package main

import (
	"fmt"
	"net"
	"os"

	"github.com/vukomanv/blastell/internal/executor"
	"github.com/vukomanv/blastell/internal/rsepparser"
)

func main() {
	fmt.Println("Blastell started...")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("error: failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error: failed accepting connection: ", err.Error())
			break
		}

		go handleClient(conn)
	}

	l.Close()
}

func handleClient(c net.Conn) {
	buffer := make([]byte, 1024)
	defer c.Close()
	for {
		n, err := c.Read(buffer)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if n == 0 {
			fmt.Println("error: input is empty")
			continue
		}

		input := buffer[:n]

		parsedInput, err := rsepparser.Parse(input)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = executor.Execute()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(parsedInput)
		// process the command + args

		c.Write([]byte("success\n"))
	}
}
