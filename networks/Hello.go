package networks

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func Listener(address string) {

	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	defer listen.Close()

	for {
		accept, err := listen.Accept()

		if err != nil {
			log.Fatal(fmt.Sprintf("accept err: %v", err))
		}
		println("accept one conn")
		//异步 处理
		go Handle(accept)
	}
}

func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		readString, err := reader.ReadString('\n')

		if err != nil {

			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}

		println(readString)
		bytes := []byte("hello " + readString)
		_, _ = conn.Write(bytes)
	}

}
