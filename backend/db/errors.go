// Package db (errors) contains Database related errors.
package db

import "errors"

var (
	// ErrRecordNotFound is returned for DB error, record not found.
	ErrRecordNotFound = errors.New("record not found")

	// ErrDuplicateRecord is returned for DB error, record already exists.
	ErrDuplicateRecord = errors.New("duplicate record")
)
