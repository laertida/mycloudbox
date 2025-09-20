package endpoints

import (
	"log"
	"path/filepath"
)

type Producer interface {
	Configure()
}

type File struct {
	Protocol   string
	Path       string
	Properties string
}

func (f *File) Log() {
	log.Printf("protocol: %s path: %s, properties: %s\n", f.Protocol, f.Path, f.Properties)
	path := filepath.Join(f.Path)
	log.Printf("dir(p): %s, %v\n", filepath.Dir(path), filepath.IsAbs(f.Path))
}
