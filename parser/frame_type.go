package parser

// FrameType type
type FrameType byte

const (
	// FrameString identifies a string frame.
	FrameString FrameType = iota

	// FrameBinary identifies a binary frame.
	FrameBinary
)
