package main

import (
	"chat/tcpip"
	"fmt"
	"log"
	"os"
)

func HandleArgs() {
	argsWithProgram := os.Args[1:]

	fmt.Println(argsWithProgram)

	if len(argsWithProgram) > 0 {
		switch argsWithProgram[0] {
		case "-server":
			tcpip.Listen("tcp", ":5032")
			break
		case "-client":
			tcpip.Connect("tcp", "localhost:5032")
			break
		}
	}
}

func main() {
	log.Println("Start Programs..")

	HandleArgs()
}
