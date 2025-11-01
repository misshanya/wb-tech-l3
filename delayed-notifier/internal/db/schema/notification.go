package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Time("scheduled_at"),

		field.String("title").
			NotEmpty(),

		field.String("content").
			NotEmpty(),

		field.String("channel").
			NotEmpty(),

		field.String("receiver").
			NotEmpty(),

		field.String("status").
			NotEmpty(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return nil
}
