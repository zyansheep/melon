package packet

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"

	"../../util"
)

// Packet represents a packet to be sent or recieved by Melon or the client.
type Packet interface {
	ID() int32
	Read(reader io.Reader) error
	Write(writer io.Writer) error
}

// ReadPacket reads a single packet from the provided connection.
func ReadPacket(conn net.Conn) (Packet, error) {
	reader := bufio.NewReader(conn)

	_, err := util.ReadVarint(reader)
	if err != nil {
		return nil, err
	}

	id, err := util.ReadVarint(reader)
	if err != nil {
		return nil, err
	}

	pk := Packet(nil)

	switch id {
	case HandshakePacketID:
		pk = new(HandshakePacket)
		err = pk.Read(reader)
	case PingPacketID:
		pk = new(PingPacket)
		err = pk.Read(reader)
	default:
		fmt.Printf("Recieved unknown packet ID: %v.\n", id)
	}

	if err != nil {
		return nil, err
	}

	return pk, nil
}

// WritePacket writes a single packet to the provided connection.
func WritePacket(conn net.Conn, pk Packet) error {
	writer := bufio.NewWriter(conn)

	tempBuffer := bytes.Buffer{}
	tempWriter := bufio.NewWriter(&tempBuffer)

	err := util.WriteVarint(pk.ID(), tempWriter)
	if err != nil {
		return err
	}

	err = pk.Write(tempWriter)
	if err != nil {
		return err
	}

	err = tempWriter.Flush()
	if err != nil {
		return err
	}

	err = util.WriteVarint(int32(tempBuffer.Len()), writer)
	if err != nil {
		return err
	}

	_, err = writer.Write(tempBuffer.Bytes())
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
