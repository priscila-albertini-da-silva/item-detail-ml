package repository

import (
	"errors"
	"io"
)

type mockFileOpener struct {
	reader io.Reader
	err    error
}

func (m *mockFileOpener) Open(name string) (io.ReadCloser, error) {
	if m.err != nil {
		return nil, m.err
	}
	return io.NopCloser(m.reader), nil
}

type errorReader struct{}

func (errorReader) Read(p []byte) (int, error) {
	return 0, errors.New("simulated read error")
}
