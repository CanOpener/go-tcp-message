package gtm

import (
	"log"
	"net"
)

const (
	WriterStopped = iota
	ReaderStopped
	ConnectionClosed
	WriterStarted
	ReaderStarted
)

type conn struct {
	Netcon            net.Conn
	IncommingMessages chan *[]byte
	InfoChan          chan int
	internalComsChan  chan int
	outgoingMessages  chan *[]byte
	writerKiller      chan bool
	MessagesSent      uint64
	MessagesReceived  uint64
	ReaderListening   bool
	WriterListening   bool
}

func NewConnection(con net.Conn) conn {
	return conn{
		Netcon:            con,
		IncommingMessages: make(chan *[]byte, 100),
		InfoChan:          make(chan int, 5),
		internalComsChan:  make(chan int, 5),
		outgoingMessages:  make(chan *[]byte, 100),
		writerKiller:      make(chan bool, 1),
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

func (c *conn) Close() {
	err := c.Netcon.Close()
	if err != nil {
		fatalLog(err)
	}
	c.killWriter()
}
