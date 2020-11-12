package tcpip

import (
	"fmt"
	"io"
	"log"
	"net"
)

var conn net.Conn
var err error

func SendMessage(message string) {
	if conn != nil {
		conn.Write([]byte(message))
	}
}

func Input() {
	var message string

	for {
		fmt.Print(">")
		fmt.Scanln(&message)

		SendMessage(message)
		fmt.Println()
	}
}

func Read() {
	buf := make([]byte, 1024) //1kb
	fmt.Println("Read from servers...")

	for {
		count, error := conn.Read(buf)

		if nil != error {
			if io.EOF == error {
				log.Printf("connection is closed from server; %v", conn.RemoteAddr().String())
			}

			log.Printf("fail to receive data; err: %v", error)
			return
		}

		if count > 0 {
			data := buf[:count]
			fmt.Println("message> " + string(data))
		}
	}
}

func Connect(protocol string, host string) {
	conn, err = net.Dial(protocol, host)

	if nil != err {
		log.Fatalf("failed to connect to server err: %v", err)
	}
	defer conn.Close()

	go Read()
	Input()
}
