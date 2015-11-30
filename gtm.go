package gtm

import (
	"log"
	"net"
)

type conn struct {
	netcon            net.Conn
	IncommingMessages chan *[]byte
	InfoChan          chan int
	internalComsChan  chan int
	outgoingMessages  chan *[]byte
	MessagesSent      int
	MessagesReceived  int
	ReaderListening   bool
	WriterListening   bool
}

func NewConnection(con net.Conn) conn {
	return conn{
		netcon:            con,
		IncommingMessages: make(chan *[]byte, 100),
		InfoChan:          make(chan int, 5),
		internalComsChan:  make(chan int, 5),
		outgoingMessages:  make(chan *[]byte, 100),
		MessagesSent:      0,
		MessagesReceived:  0,
		ReaderListening:   false,
		WriterListening:   false,
	}
}

func fatalLog(v ...interface{}) {
	// later this will use a custom first class
	// logging function specified by the user
	log.Fatalln(v)
}