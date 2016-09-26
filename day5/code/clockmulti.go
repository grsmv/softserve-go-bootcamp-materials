package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

const (
	listenDefault = "localhost:8000"
)

var bindTo string

func init() {
	flag.StringVar(&bindTo, "listen", listenDefault, "Define listen host default is: "+listenDefault)
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", bindTo)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	timeout := time.After(5 * time.Second)
	for {
		select {
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
			if err != nil {
				return // e.g., client disconnected
			}
			time.Sleep(1 * time.Second)

		case <-timeout:
			return
		}
	}
}
