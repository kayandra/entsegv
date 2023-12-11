package schema

import (
	"entsegv/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct{ ent.Schema }

// Fields of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.CreateBaseMixin("usr",
			field.String("email").Unique().NotEmpty(),
			field.String("name").Optional()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("apps", Application.Type).Through("membership", Member.Type),
	}
}
