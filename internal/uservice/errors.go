package uservice

import "emperror.dev/errors"

var ErrAlreadyExists = errors.New("entity already exists")

var ErrNotFound = errors.New("entity not found")

var ErrInvalidFieldValue = errors.New("field value is invalid")
