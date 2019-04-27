package main

import (
	"github.com/google/uuid"
)

type Image struct {
	ID      *uuid.UUID `json:"uuid,omitempty"`
	Name    string     `json:"name,omitempty"`
	Version string     `json:"version,omitempty"`
}
