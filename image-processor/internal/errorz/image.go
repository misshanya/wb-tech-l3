package errorz

import "errors"

var (
	ErrImageNotFound  = errors.New("image not found")
	ErrImageIsNotDone = errors.New("image is not done")
)
