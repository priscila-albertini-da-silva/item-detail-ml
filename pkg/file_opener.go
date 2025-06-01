package pkg

import (
	"io"
	"os"
)

type FileOpener interface {
	Open(name string) (io.ReadCloser, error)
}

type OsFileOpener struct{}

func (OsFileOpener) Open(name string) (io.ReadCloser, error) {
	return os.Open(name)
}
