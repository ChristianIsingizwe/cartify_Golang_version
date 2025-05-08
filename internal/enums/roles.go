package enums

import (
	"database/sql/driver"
	"fmt"
)

type UserRoles string

const (
	Admin  UserRoles = "admin"
	Seller UserRoles = "user"
	Buyer  UserRoles = "buyer"
)

func (r *UserRoles) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan role: %T", value)
	}
	*r = UserRoles(string(b))
	return nil
}

func (r UserRoles) Value() (driver.Value, error) {
	return string(r), nil
}
