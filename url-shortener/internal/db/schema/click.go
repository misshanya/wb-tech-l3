package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Click holds the schema definition for the Click entity.
type Click struct {
	ent.Schema
}

// Fields of the Click.
func (Click) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("user_agent"),
		field.String("ip"),

		field.Time("clicked_at"),
	}
}

// Edges of the Click.
func (Click) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("link", Link.Type).
			Ref("clicks").
			Unique(),
	}
}
