package adrgo

import (
	"errors"
)

var (
	ErrAlreadyExists   = errors.New("already exists")
	ErrWriteFileFailed = errors.New("write file failed")
)
