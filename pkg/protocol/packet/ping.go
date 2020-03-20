package packet

import (
	"encoding/binary"
	"io"
)

// PingPacket ...
type PingPacket struct {
	Payload int64
}

// ID ...
func (pk PingPacket) ID() int32 {
	return PingPacketID
}

// Read ...
func (pk *PingPacket) Read(reader io.Reader) error {
	err := binary.Read(reader, binary.BigEndian, &pk.Payload)
	if err != nil {
		return err
	}

	return nil
}

// Write ...
// PingPacket should not be written by the server.
func (pk PingPacket) Write(writer io.Writer) error {
	return nil
}
