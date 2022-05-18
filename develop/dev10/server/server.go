package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type server struct {
	conn net.Conn
}

func newServer() *server {
	return &server{}
}

func (srv *server) startServer() {

	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Test server started...")

	srv.conn, err = ln.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		message, err := bufio.NewReader(srv.conn).ReadString('\n')
		if err != nil && err != io.EOF {
			//log.Println(err)
			// return if closed conn, no need error
			break
		}
		answer := strings.ToUpper(message)
		answer = "rewrite: " + answer
		_, _ = srv.conn.Write([]byte(answer))
		//if err != nil {
		//log.Println(err)
		// return if closed conn, no need error
		//}
	}
}

func (srv *server) writeString(s string) {
	_, _ = srv.conn.Write([]byte(s))
}

func main() {
	serverOk := newServer()
	go serverOk.startServer()
	time.Sleep(time.Second*15)
	serverOk.writeString("l2l2l2")
	time.Sleep(time.Second)
	for i:=1; i<10; i++ {
	 	serverOk.writeString("l2l2l2")
	 	time.Sleep(time.Second)
	}
	time.Sleep(time.Second*10)
}