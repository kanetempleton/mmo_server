package net

import (
	"encoding/binary"
	"fmt"
)

const PACKET_ID_MESSAGE = 1

// KProtocol for client-server communication
type KProtocol struct {
	version [4]byte;
}

// NewConnectionManager creates a new ConnectionManager instance.
func NewKProtocol() *KProtocol {
	return &KProtocol{
		
	}
}

func (kp *KProtocol) ProcessPacket(dat []byte) {
	for i := 0; i < 1024; i++ {
		fmt.Printf("%d\t%c\t//%d\n", i, dat[i], dat[i])
	}

	//packetVersion := dat[0:4]
	//packetKey := dat[4:20]

	if (len(dat)==1024) {
	packetID := int(dat[20])<<8 | int(dat[21])
	payloadLength := int(dat[22])<<8 | int(dat[23])

	packetPayload := dat[24 : 24+payloadLength]

	fmt.Printf("packet processed: %d len=%d\npayload:%s\n", packetID, payloadLength, packetPayload)
	} else {
		fmt.Printf("Malformed packet, len=%d\n",len(dat));
	}
}


// SendMessagePacket sends a console log message to the client
func (kp *KProtocol) SendMessagePacket(msg string) *KPacket {
	// Convert the message to a byte slice
	msgBytes := []byte(msg)

	// Get the length of the message
	payloadLength := len(msgBytes)

	// Create a new packet with MESSAGE_PACKET_ID
	packet := NewKPacket(PACKET_ID_MESSAGE, payloadLength, msgBytes)
	fmt.Printf("New Packet %d %d\n",packet.packetID,packet.payloadLength)
	//packet.ConstructPacket(PACKET_ID_MESSAGE, payloadLength, msgBytes)

	return packet
}

// ConstructPacketBytes creates a packet with specified packetid, payloadLength, and dat
func (kp *KProtocol) ConstructPacketBytes(packetid, payloadLength int, dat []byte) []byte {
	ret := make([]byte, 1024)

	// Construct packet version (4 bytes)
	copy(ret[0:4], kp.version[:])

	// Client-key (16 bytes)
	// No specific initialization for now, you can customize as needed

	// Packet-id (2 bytes)
	binary.BigEndian.PutUint16(ret[4:6], uint16(packetid))

	// Payload-length (2 bytes)
	binary.BigEndian.PutUint16(ret[6:8], uint16(payloadLength))

	// Payload
	copy(ret[8:], dat)

	return ret
}