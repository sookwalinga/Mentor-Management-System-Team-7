package db

import "errors"

var (
	ErrRecordNotFound  = errors.New("record not found") // requested record is not found.
	ErrDuplicateRecord = errors.New("duplicate record") // requested record already exists.
)
