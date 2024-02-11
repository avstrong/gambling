package gmanager

import "emperror.dev/errors"

var ErrInvalidFieldValue = errors.New("field value is invalid")

var ErrAlreadyExists = errors.New("already exists")
