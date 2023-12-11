package schema

import (
	"entsegv/mixins"
	"net/http"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	AttemptStatusSuccess   = "success"
	AttemptStatusMalformed = "malformed"
	AttemptStatusTimeout   = "timeout"
	AttemptStatusFailed    = "failed"
)

// Attempt holds the schema definition for the Attempt entity.
type Attempt struct{ ent.Schema }

// Fields of the Attempt.
func (Attempt) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.CreateBaseMixin("attempt",
			field.String("response").Default("{}"),
			field.Int("status_code").Default(http.StatusOK).
				StructTag(`json:"status_code"`),
			field.Int("duration_ms").Default(0).
				StructTag(`json:"duration_ms"`),
			field.String("application_id").Immutable(),
			field.String("message_id").Immutable(),
			field.Enum("status").Values(
				AttemptStatusSuccess,
				AttemptStatusMalformed,
				AttemptStatusTimeout,
				AttemptStatusFailed,
			).Default(AttemptStatusSuccess),
		),
	}
}

// Edges of the Attempt.
func (Attempt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message", Message.Type).
			Ref("attempts").
			Field("message_id").
			Immutable().
			Required().
			Unique(),
		edge.From("application", Application.Type).
			Ref("attempts").
			Field("application_id").
			Immutable().
			Required().
			Unique(),
	}
}
