package schema

import (
	"entsegv/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Application holds the schema definition for the Application entity.
type Application struct{ ent.Schema }

// Fields of the Application.
func (Application) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.CreateBaseMixin("app", field.String("name")),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", User.Type).Ref("apps").Through("membership", Member.Type),
		edge.To("consumers", Consumer.Type),
		edge.To("messages", Message.Type),
		edge.To("attempts", Attempt.Type),
	}
}
