package schema

import (
	"time"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.String("description"),
		field.Bool("is_done").
			Annotations(entoas.Annotation{
				ReadOnly: true,
			}).
			Optional().
			Default(false),
		field.Time("created_at").
			Annotations(entoas.Annotation{
				ReadOnly: true,
			}).
			Default(time.Now),
		field.Time("updated_at").
			Annotations(entoas.Annotation{
				ReadOnly: true,
			}).
			Default(time.Now()),
		field.String("assigned_id").
			Annotations(entoas.Annotation{
				ReadOnly: true,
			}).
			Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assigned", User.Type).
			Ref("tasks").
			Field("assigned_id").
			Unique().
			Annotations(entoas.Annotation{
				ReadOnly: true,
			}),
	}
}
