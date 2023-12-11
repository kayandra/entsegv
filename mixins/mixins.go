package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/rs/xid"
)

type BaseMixin struct {
	mixin.Schema
	idprefix string
	fields   []ent.Field
}

func (b BaseMixin) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return b.idprefix + "_" + xid.New().String()
			}).
			Immutable().
			Unique().
			NotEmpty(),
	}
	fields = append(fields, b.fields...)
	return append(fields, []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}...)
}

func CreateBaseMixin(idprefix string, fields ...ent.Field) BaseMixin {
	return BaseMixin{idprefix: idprefix, fields: fields}
}
