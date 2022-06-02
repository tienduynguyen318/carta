package main

import (
	"encoding/json"
	"io"
	"log"
)

type outputWriter struct {
	logger *log.Logger
}

func NewOutputWriter(writer io.Writer) (*outputWriter, error) {
	logger := log.New(writer, "", 0)
	return &outputWriter{logger: logger}, nil
}

func (ow *outputWriter) Print(input interface{}) {
	b, err := json.MarshalIndent(input, "", "  ")
	if err == nil {
		ow.logger.Println(string(b))
	}
	return
}
