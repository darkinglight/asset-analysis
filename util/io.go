package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const (
	IncomeStorage = "income.store"
)

type Storage interface {
	Open(filePath string) error
	Close() error
	Write(v interface{}) error
	Read(v interface{}) error
}

type storage struct {
	file *os.File
	io   *bufio.ReadWriter
}

func (s *storage) Open(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	s.file = file
	s.io = bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file))
	return nil
}

func (s *storage) Close() error {
	s.io.Flush()
	s.file.Close()
	return nil
}

func (s *storage) Write(v interface{}) error {
	str, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprintln(*s.io, str)
	return nil
}
