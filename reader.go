package gtm

import (
	"bytes"
	"encoding/binary"
)

func (c *conn) startReader() {
	c.setReaderListening(true)
	defer c.setReaderListening(false)

	var messageBuffer bytes.Buffer
	var bytesToRead int

	for {
		buf := make([]byte, 1400)

		dataSize, err := c.netcon.Read(buf)
		if err != nil {
			// connection closed
			return
		}
		data := buf[0:dataSize]
		messageBuffer.Write(data)

		if messageBuffer.Len() >= bytesToRead {
			for {
				if bytesToRead != 0 {
					message := make([]byte, bytesToRead)
					messageBytes, err := messageBuffer.Read(message)
					if err != nil {
						fatalLog(err)
					}
					if messageBytes != bytesToRead {
						fatalLog("Something went wrong, bytes to read != read bytes")
					}
					c.IncommingMessages <- &message
					c.MessagesReceived++
					bytesToRead = 0
				}

				if bytesToRead == 0 && messageBuffer.Len() > 2 {
					btrBuffer := make([]byte, 2)
					btrBytes, err := messageBuffer.Read(btrBuffer)
					if err != nil {
						fatalLog(err)
					}
					if btrBytes != 2 {
						fatalLog("Something went wrong, btrBytes != 2")
					}

					bytesToRead = int(binary.LittleEndian.Uint16(btrBuffer))
				}

				if bytesToRead == 0 || (messageBuffer.Len() < bytesToRead) {
					break
				}
			}
		}
	}
}

func (c *conn) setReaderListening(listening bool) {
	c.ReaderListening = listening
}
