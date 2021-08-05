package uuid

import (
	"database/sql/driver"
	"fmt"
)

// Value implements the driver.Valuer interface.
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}

// Scan implements the sql.Scanner interface.
// A 16-byte slice is handled by UnmarshalBinary, while
// a longer byte slice or a string is handled by UnmarshalText.
func (u *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		if len(src) == Size {
			return u.UnmarshalBinary(src)
		}
		return u.UnmarshalText(src)

	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("uuid: cannot convert %T to UUID", src)
}

// NullUUID can be used with the standard sql package to represent a
// UUID value that can be NULL in the database
type NullUUID struct {
	UUID  UUID
	Valid bool
}

// Value implements the driver.Valuer interface.
func (u NullUUID) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	// Delegate to UUID Value function
	return u.UUID.Value()
}

// Scan implements the sql.Scanner interface.
func (u *NullUUID) Scan(src interface{}) error {
	if src == nil {
		u.UUID, u.Valid = Nil, false
		return nil
	}

	// Delegate to UUID Scan function
	u.Valid = true
	return u.UUID.Scan(src)
}
