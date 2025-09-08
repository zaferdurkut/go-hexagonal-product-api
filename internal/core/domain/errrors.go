package domain

import "errors"

var (
	ErrProductNameInvalid = errors.New("product name is invalid")
	ErrProductNotFound    = errors.New("product not found")
)
