package tcpip

import (
	"io"
	"log"
	"net"
)

var sessions [100]net.Conn
var offset = 0

func Handle(connection net.Conn) {
	//세션에 추가함
	sessions[offset] = connection
	offset++
	buf := make([]byte, 1024) //1kb

	for {
		count, error := connection.Read(buf)

		if nil != error {
			if io.EOF == error {
				log.Printf("connection is closed from client; %v", connection.RemoteAddr().String())
			}

			log.Printf("fail to receive data; err: %v", error)
			return
		}

		var data []byte
		if count > 0 {
			data = buf[:count]
		}

		//Broadcast to sessions
		for i := 0; i < offset; i++ {
			client := sessions[i]
			client.Write(data)
		}
	}
}

func Listen(protocol string, port string) {
	listener, error := net.Listen(protocol, port)
	log.Println("Listening Clients...")

	if nil != error {
		log.Fatalf("fail to bind address to 10000 err: %v", error)
	}
	defer listener.Close()

	log.Println("Listening Clients...")
	for {
		conn, error := listener.Accept()

		if nil != error {
			log.Printf("fail to accept; err: %v", error)
			continue
		}

		go Handle(conn)
	}
}
