package main

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid"
)

type ID interface {		
	MarshalText() ([] byte, error)
}
