package packet

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"reflect"

	"../../util"
)

// Packet represents a packet to be sent or recieved by Melon or the client.
type Packet struct{
	ID int32
	data []byte
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
	var data []byte;
	reader.Read(data);
	
	return Packet{id, data}
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
