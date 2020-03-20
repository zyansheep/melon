package packet

import (
	"encoding/binary"
	"io"

	"github.com/melondevs/melon/pkg/util"
)

// HandshakePacket ...
type HandshakePacket struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

// ID ...
func (pk HandshakePacket) ID() int32 {
	return HandshakePacketID
}

// Read ...
func (pk *HandshakePacket) Read(reader io.Reader) error {
	var err error

	pk.ProtocolVersion, err = util.ReadVarint(reader)
	if err != nil {
		return err
	}

	pk.ServerAddress, err = util.ReadString(reader)
	if err != nil {
		return err
	}

	err = binary.Read(reader, binary.BigEndian, &pk.ServerPort)
	if err != nil {
		return err
	}

	pk.NextState, err = util.ReadVarint(reader)
	if err != nil {
		return err
	}

	return nil
}

// Write ...
// HandshakePacket should not be written by the server.
func (pk HandshakePacket) Write(writer io.Writer) error {
	return nil
}
