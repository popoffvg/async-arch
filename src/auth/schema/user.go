package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	gen "github.com/popoffvg/async-arch/auth/internal/ent"
	"github.com/popoffvg/async-arch/auth/internal/ent/hook"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty(),
		field.String("login").
			NotEmpty().
			Unique(),
		field.Bytes("password").
			NotEmpty(),
		field.Enum("scopes").
			Values("USER", "ADMIN", "SUPERUSER").
			Default("USER"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (gen.Value, error) {
					m.SetID(uuid.New().String())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate),
	}
}
