package main

import (
	"io/ioutil"
)

type File struct {
}

func (f File) read(name string) (string, error) {
	bytesContent, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(bytesContent), nil
}
