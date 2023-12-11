package schema

import (
	"entsegv/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	RoleOwner   = "owner"
	RoleManager = "manager"
	RoleMember  = "member"
)

// Member holds the schema definition for the Member entity.
type Member struct{ ent.Schema }

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "application_id"),
	}
}

// Fields of the Member.
func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.CreateBaseMixin("mem",
			field.String("user_id"),
			field.String("application_id"),
			field.Enum("role").Values(
				RoleOwner,
				RoleManager,
				RoleMember,
			).Default(RoleMember)),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required().Unique().Field("user_id"),
		edge.To("application", Application.Type).Required().Unique().Field("application_id"),
	}
}
