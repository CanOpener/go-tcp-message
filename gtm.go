package gtm

import (
	"log"
	"net"
)

type conn struct {
	netcon            net.Conn
	IncommingMessages chan *[]byte
	InfoChan          chan int
	outgoingMessages  chan *[]byte
	MessagesSent      int
	MessagesReceived  int
	ReaderListening   bool
	WriterListening   bool
}
