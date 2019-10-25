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
	Create() error
	Open() error
	Close() error
	Write(v interface{}) error
	Read() ([]byte, error)
}

type storage struct {
	filePath string
	file     *os.File
	io       *bufio.ReadWriter
}

func (s *storage) Create() error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	s.file = file
	s.io = bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file))
	return nil
}

func (s *storage) Open() error {
	file, err := os.Open(s.filePath)
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
	fmt.Fprintln(*s.io, string(str))
	return nil
}

func (s *storage) Read() ([]byte, error) {
	v, err := s.io.ReadBytes('\n')
	return v, err
}

func NewStore(filePath string, recreate bool) (Storage, error) {
	s := &storage{filePath: filePath}
	var err error
	if recreate {
		err = s.Create()
	} else {
		err = s.Open()
	}
	return s, err
}
