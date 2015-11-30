package gtm

import (
	"encoding/binary"
)

func (c *conn) startWriter() {
	c.setWriterListening(true)
	defer c.setWriterListening(false)

	for {
		select {
		case m := <-c.outgoingMessages:
			length := uint16(len(*m))
			lenBytes := make([]byte, 2)
			binary.LittleEndian.PutUint16(lenBytes, length)
			_, err := c.netcon.Write(append(lenBytes, *m...))
			if err != nil {
				fatalLog(err)
			}
		case <-c.writerKiller:
			return
		}
	}
}

func (c *conn) setWriterListening(listening bool) {
	c.WriterListening = listening
}

func (c *conn) killWriter() {
	c.writerKiller <- false
}
