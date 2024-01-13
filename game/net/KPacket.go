package net

import (
	"encoding/binary"
)

// KPacket represents a packet in the KProtocol.
type KPacket struct {
	version      [4]byte
	clientKey    [16]byte
	packetID     uint16
	payloadLength uint16
	payload      []byte
}

// NewKPacket creates a new KPacket instance.
func NewKPacket(packetID, payloadLength int, payload []byte) *KPacket {
	return &KPacket{
		// Set version, clientKey as needed
		packetID:     uint16(packetID),
		payloadLength: uint16(payloadLength),
		payload:      payload,
	}
}

// ConstructPacket constructs a byte slice representing the KPacket.
func (kp *KPacket) PacketBytes() []byte {
	ret := make([]byte, 1024)

	// Construct packet version (4 bytes)
	copy(ret[0:4], kp.version[:])

	// Client-key (16 bytes)
	copy(ret[4:20], kp.clientKey[:])

	// Packet-id (2 bytes)
	binary.BigEndian.PutUint16(ret[20:22], kp.packetID)

	// Payload-length (2 bytes)
	binary.BigEndian.PutUint16(ret[22:24], kp.payloadLength)

	// Payload
	copy(ret[24:], kp.payload)

	return ret
}

func (kp* KPacket) PayloadLength() uint16 {
	return kp.payloadLength;
}
