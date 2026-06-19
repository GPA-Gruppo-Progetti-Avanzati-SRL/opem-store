package sql

import (
	stdsql "database/sql"
	"errors"
)

func isNotFound(err error) bool {
	return errors.Is(err, stdsql.ErrNoRows)
}
