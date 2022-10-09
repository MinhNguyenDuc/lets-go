package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no mathcing record found")

type Snippet struct {
	ID       int
	Title    string
	Content  string
	Created  time.Time
	Exprires time.Time
}
