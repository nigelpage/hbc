package store

import "errors"

// Sentinel errors indicate that a certain, often expected, category of error has occurred.

var ErrNotFound = errors.New("Record not found")
