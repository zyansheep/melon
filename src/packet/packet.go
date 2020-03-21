package packet

import (
	"bufio"
	"bytes"
	_ "io"
	"net"
	_ "reflect"

	"../util"
	_ "../protocol"
)

// Packet represents a packet to be sent or recieved by Melon or the client.
type Packet struct{
	ID int32
	Data []byte
}

// ReadPacket reads a single packet from the provided connection.
func ReadPacket(conn net.Conn) (*Packet, error) {
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
	
	return &Packet{id, data}, nil
}

// WritePacket writes a single packet to the provided connection.
func WritePacket(conn net.Conn, pk Packet) error {
	writer := bufio.NewWriter(conn)

	tempBuffer := bytes.Buffer{}
	tempWriter := bufio.NewWriter(&tempBuffer)

	//Write byte ID
	err := util.WriteVarint(pk.ID, tempWriter)
	if err != nil {return err}
	
	//Write packet data
	_, err = tempWriter.Write(pk.Data);
	if err != nil {return err}
	
	//Flush data in tempbuffer
	tempWriter.Flush()
	
	//Write length to real buffer
	err = util.WriteVarint(int32(tempBuffer.Len()), writer)
	if err != nil {return err}
	
	//Write tempbuffer
	_, err = writer.Write(tempBuffer.Bytes())
	if err != nil {return err}
	
	//Flush buffer to client
	err = writer.Flush()
	if err != nil {return err}

	return nil
}
