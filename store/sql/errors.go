package sql

import "errors"

// ErrNotSupported is returned by repository methods not implemented for the SQL driver.
var ErrNotSupported = errors.New("operation not supported by sql store")
