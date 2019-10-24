package util

import (
	"testing"
)

func TestOpenFile(t *testing.T) {
	s := &storage{}
	err := s.Open(IncomeStorage)
	if err != nil {
		t.Error("open file failed:", err)
	}
	s.Close()
}

func TestWriteStruct(t *testing.T) {
	s := &storage{}
	_ = s.Open(IncomeStorage)
	data := struct {
		Field1 string
	}{"test"}
	s.Write(data)
	s.Close()
}
