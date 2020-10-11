package parser

import "io"

type writer struct{}

func (writer) Write(p []byte) (int, error) {
	for _, b := range p {
		HandleByte(b)
	}

	return len(p), nil
}

func MakeWriter() io.Writer {
	return writer{}
}
