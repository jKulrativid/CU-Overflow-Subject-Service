package entity

import "errors"

var ErrUnauthorized = errors.New("unauthorized")
var ErrNotFound = errors.New("not found")
var ErrConstraintViolation = errors.New("violating data integrity")
