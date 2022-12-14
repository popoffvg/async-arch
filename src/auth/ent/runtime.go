// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/popoffvg/async-arch/auth/ent/schema"
	"github.com/popoffvg/async-arch/auth/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescLogin is the schema descriptor for login field.
	userDescLogin := userFields[1].Descriptor()
	// user.LoginValidator is a validator for the "login" field. It is called by the builders before save.
	user.LoginValidator = userDescLogin.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func([]byte) error)
	// userDescScopes is the schema descriptor for scopes field.
	userDescScopes := userFields[3].Descriptor()
	// user.DefaultScopes holds the default value on creation for the scopes field.
	user.DefaultScopes = schema.Scope(userDescScopes.Default.(string))
	// user.ScopesValidator is a validator for the "scopes" field. It is called by the builders before save.
	user.ScopesValidator = func() func(string) error {
		validators := userDescScopes.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(scopes string) error {
			for _, fn := range fns {
				if err := fn(scopes); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(string)
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
