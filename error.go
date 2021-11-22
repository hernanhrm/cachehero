package cachehero

import "errors"

var (
	ErrNotFound = errors.New("key not found")
	ErrNotEnteredKeys = errors.New("you must enter at least one key")
)
