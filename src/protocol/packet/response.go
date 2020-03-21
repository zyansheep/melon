package packet

import (
	"io"

	"../../util"
)

// ResponsePacket ...
type ResponsePacket struct {
	Response string
}

// ID ...
func (pk ResponsePacket) ID() int32 {
	return ResponsePacketID
}

// Read ...
// ResponsePacket should not be read by the server.
func (pk *ResponsePacket) Read(reader io.Reader) error {
	return nil
}

// Write ...
func (pk ResponsePacket) Write(writer io.Writer) error {
	err := util.WriteString(pk.Response, writer)
	if err != nil {
		return err
	}

	return nil
}
