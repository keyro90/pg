package pg

import (
	"net"

	"github.com/keyro90/pg/v10/internal"
)

// ErrNoRows is returned by QueryOne and ExecOne when query returned zero rows
// but at least one row is expected.
var ErrNoRows = internal.ErrNoRows

// ErrMultiRows is returned by QueryOne and ExecOne when query returned
// multiple rows but exactly one row is expected.
var ErrMultiRows = internal.ErrMultiRows

// Error represents an error returned by PostgreSQL server
// using PostgreSQL ErrorResponse protocol.
//
// https://www.postgresql.org/docs/10/static/protocol-message-formats.html
type Error interface {
	error

	// Field returns a string value associated with an error field.
	//
	// https://www.postgresql.org/docs/10/static/protocol-error-fields.html
	Field(field byte) string

	// IntegrityViolation reports whether an error is a part of
	// Integrity Constraint Violation class of errors.
	//
	// https://www.postgresql.org/docs/10/static/errcodes-appendix.html
	IntegrityViolation() bool
}

var _ Error = (*internal.PGError)(nil)

func isBadConn(err error, allowTimeout bool) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(internal.Error); ok {
		return false
	}
	if pgErr, ok := err.(Error); ok {
		return pgErr.Field('S') == "FATAL"
	}
	if allowTimeout {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return !netErr.Temporary()
		}
	}
	return true
}

//------------------------------------------------------------------------------

type timeoutError interface {
	Timeout() bool
}
