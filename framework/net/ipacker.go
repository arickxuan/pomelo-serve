package net

import "io"

type IPacker interface {
	Pack(message *ProtoMessage) ([]byte, error)
	Unpack(reader io.Reader) (*ProtoMessage, error)
}
