package schema

import (
	"entsegv/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

const (
	MessageStatusProcessing = "processing"
	MessageStatusSuccess    = "success"
	MessageStatusFailed     = "failed"
)

// Message holds the schema definition for the Message entity.
type Message struct{ ent.Schema }

// Fields of the Message.
func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.CreateBaseMixin("msg",
			field.String("uid").Nillable().Optional(),
			field.String("endpoint").NotEmpty(),
			field.String("payload").Default("{}"),
			field.String("application_id").Immutable(),
			field.String("consumer_id").Immutable(),
			field.Int("attempt_count").Default(0),
			field.Time("last_attempt_at").Nillable().Optional(),
			field.Time("next_attempt_at").Nillable().Optional(),
			field.Enum("status").Values(
				MessageStatusProcessing,
				MessageStatusSuccess,
				MessageStatusFailed,
			).Default(MessageStatusProcessing),
		),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("consumer", Consumer.Type).
			Ref("messages").
			Field("consumer_id").
			Immutable().
			Required().
			Unique(),
		edge.From("application", Application.Type).
			Ref("messages").
			Field("application_id").
			Immutable().
			Required().
			Unique(),
		edge.To("attempts", Attempt.Type),
	}
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uid").Edges("consumer").Unique(),
	}
}
