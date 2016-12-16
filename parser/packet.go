package parser

// PacketType
type PacketType int

const (

	// OPEN is sent from the server when a new transport is opened (recheck).
	OPEN PacketType = iota

	// CLOSE is conn the close of this transport but does not shutdown
	CLOSE

	// PONG is sent by the server to respond to ping packets.
	PING

	// PONG is sent by the server to respond to ping packets.
	PONG

	// MESSAGE is actual message, client and server should call their callbacks
	// with the data.
	MESSAGE

	// UPGRADE is sent before engine.io switches a transport to test if server
	// and client can communicate over this transport. If this test succeed,
	// the client sends an upgrade packets which requests the server to flush
	// its cache on the old transport and switch to the new transport.
	UPGRADE

	// NOOP is a noop packet. Used primarily to force a poll cycle when an
	// incoming websocket connection is received.
	NOOP
)

// convert PacketType to string
func (id PacketType) String() string {
	switch id {
	case OPEN:
		return "open"
	case CLOSE:
		return "close"
	case PING:
		return "ping"
	case PONG:
		return "pong"
	case MESSAGE:
		return "message"
	case UPGRADE:
		return "upgrade"
	case NOOP:
		return "noop"
	default:
		return "unknow"
	}
}

// StringByte converts a PacketType to byte in string
func (id PacketType) StringByte() byte {
	return byte(id) + '0'
}

// BinaryByte converts a PacketType to byte in binary
func (id PacketType) BinaryByte() byte {
	return byte(id)
}
