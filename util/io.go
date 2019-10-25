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
	Read() ([]byte, error)
}

type storage struct {
	filePath string
	file     *os.File
	io       *bufio.ReadWriter
}

func (s *storage) Open() error {
	file, err := os.Create(s.filePath)
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

func (s *storage) Read() ([]byte, error) {
	v, err := s.io.ReadBytes('\n')
	return v, err
}

func NewStore(filePath string) (Storage, error) {
	s := &storage{filePath: filePath}
	err := s.Open()
	return s, err
}
