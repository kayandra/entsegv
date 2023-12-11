package schema

import (
	"entsegv/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Consumer holds the schema definition for the Consumer entity.
type Consumer struct{ ent.Schema }

// Fields of the Consumer.
func (Consumer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.CreateBaseMixin("con",
			field.String("uid").Optional(),
			field.String("name").NotEmpty(),
			field.Int("rps").Default(60),
			field.String("application_id").Immutable(),
		),
	}
}

// Edges of the Consumer.
func (Consumer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("application", Application.Type).
			Ref("consumers").
			Field("application_id").
			Immutable().
			Required().
			Unique(),
		edge.To("messages", Message.Type),
	}
}

func (Consumer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uid").Edges("application").Unique(),
	}
}
