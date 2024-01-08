package pg

import (
	"database/sql"
	"errors"
)

var (
	ErrUniqueViolation = errors.New("ErrUniqueViolation")
	ErrInternalPg      = errors.New("InternalError")
	ErrForeignKey      = errors.New("foreign key constraint")
	ErrNoRows          = sql.ErrNoRows
)
