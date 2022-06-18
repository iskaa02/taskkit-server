package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// List holds the schema definition for the List entity.
type List struct {
	ent.Schema
}

// Fields of the List.
func (List) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.Int64("theme_id"),
		field.Bool("is_deleted").Default(false),
		field.Time("last_modified").Default(time.Now()),
		field.Time("created_at").Default(time.Now()),
	}
}

func (List) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "list"},
	}
}

// Edges of the List.
func (List) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("theme", Theme.Type).
			Ref("lists").
			Unique().
			Required().
			Field("theme_id"),
		edge.To("tasks", Task.Type),
	}
}
