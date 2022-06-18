package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gopkg.in/guregu/null.v4"
)

// Theme holds the schema definition for the Theme entity.
type Theme struct {
	ent.Schema
}

// Fields of the Theme.
func (Theme) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("primary"),
		field.String("secondary").Optional().GoType(null.String{}),
	}
}

func (Theme) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "theme"},
	}
}

// Edges of the Theme.
func (Theme) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("list", List.Type),
	}
}
