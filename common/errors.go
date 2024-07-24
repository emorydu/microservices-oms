package common

import "errors"

var (
	ErrNotItems = errors.New("items must have at least one item")
)
