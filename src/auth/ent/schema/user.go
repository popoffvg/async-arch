package schema

import (
	"fmt"
	"github.com/google/uuid"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Scope string

const (
	ScopeUser      = "USER"
	ScopeAdmin     = "ADMIN"
	ScopeSuperUser = "SUPER_USER"
)

var availableScopes = []Scope{
	ScopeUser,
	ScopeAdmin,
	ScopeSuperUser,
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Default(uuid.New().String()),
		field.String("login").
			NotEmpty().
			Unique(),
		field.Bytes("password").
			NotEmpty(),
		field.String("scopes").
			NotEmpty().
			GoType(Scope("user")).
			Default("user").
			Validate(func(s string) error {
				scopes := strings.Split(s, ",")
			Scope:
				for _, s := range scopes {
					for _, v := range availableScopes {
						if v == Scope(strings.Trim(s, " ")) {
							continue Scope
						}
					}

					return fmt.Errorf("not supported scope %s. Expected: %v", s, availableScopes)
				}

				return nil
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (s Scope) And(other Scope) Scope {
	if string(s) == "" {
		return Scope(other)
	}
	return Scope(strings.Join([]string{string(s), string(other)}, ","))
}

func (s Scope) Contains(other Scope) bool {
	return strings.Contains(string(s), string(other))
}

func (s Scope) ToArray() []string {
	return strings.Split(string(s), ",")
}
