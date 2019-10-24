package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
	io *io.ReadWriter
}

func (s storage) Open(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	s.io = bufio.NewReadWriter(f, f)
}

func (s storage) Close() error {
	s.io.Flush()
	s.io.Close()
	return nil
}

func (s storage) Write(v interface{}) error {
	str, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprintln(s.io, str)
}
