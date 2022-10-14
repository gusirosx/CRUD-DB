package domain

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrEmptyAreas empty areas from blocks
var ErrEmptyAreas = errors.New("empty areas from blocks")
