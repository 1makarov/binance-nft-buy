package account

import (
	"fmt"
)

func (a *Account) HandleAccount() error {
	email,err := a.Auth.GetEmail()
	if err != nil {
		return err
	}

	fmt.Printf("The login was completed successfully: %s\n", email)
	return nil
}
