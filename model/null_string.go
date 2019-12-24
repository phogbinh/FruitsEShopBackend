package model

import (
	"database/sql"
	"strings"
)

// A NullString contains a `String` string attribute and a `Valid` boolean attribute. `Valid` is `false` if the data yielded by sql.Scan() is null, and `true` otherwise.
type NullString struct {
	sql.NullString
}

// UnmarshalJSON sets the provided NullString object to a copy of the given value, often being invoked through JSON binding.
func (thisPtr *NullString) UnmarshalJSON(value []byte) error {
	thisPtr.String = strings.Trim(string(value), `"`)
	thisPtr.Valid = true
	return nil
}
