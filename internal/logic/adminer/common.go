package adminer

import "errors"

func validateAccessValue(access string) error {
	if access == "normalAdmin" {
		return nil
	}
	if access == "superAdmin" {
		return nil
	}
	if access == "admin" {
		return nil
	}
	return errors.New("access 的值只能是: normalAdmin, superAdmin, admin")
}
