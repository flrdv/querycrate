package querycrate

import (
	"os"
	"path"
)

type File struct {
	RelPath   string
	Name      string
	Extension string
}

func (f File) Read() (data []byte, err error) {
	data, err = os.ReadFile(path.Join(f.RelPath, f.Name))

	return data, err
}

type Dir struct {
	Path  string
	Files []File
	Dirs  []Dir
}
