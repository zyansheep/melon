package packet

import (
	"encoding/binary"
	"io"
)

// PongPacket ...
type PongPacket struct {
	Payload int64
}

// ID ...
func (pk PongPacket) ID() int32 {
	return PongPacketID
}

// Read ...
// PongPacket should not be read by the server.
func (pk *PongPacket) Read(reader io.Reader) error {
	return nil
}

// Write ...
func (pk PongPacket) Write(writer io.Writer) error {
	err := binary.Write(writer, binary.BigEndian, pk.Payload)
	if err != nil {
		return err
	}

	return nil
}
