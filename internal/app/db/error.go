package db

import "github.com/pkg/errors"

var (
	ErrNotImplModel = errors.New("bean must implements [db.Model]")
)
