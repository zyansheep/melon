package protocol

// Packet IDs.
const (
	HandshakePacketID = 0 // C -> S
	RequestPacketID   = 0 // C -> S
	ResponsePacketID  = 0 // S -> C
	PingPacketID      = 1 // C -> S
	PongPacketID      = 1 // S -> C
)