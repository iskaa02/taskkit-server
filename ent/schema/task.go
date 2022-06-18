package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/iskaa02/taskkit-server/types"
	"gopkg.in/guregu/null.v4"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("list_id"),
		field.Text("description").Optional().GoType(null.String{}),
		field.Time("reminder").Optional().GoType(null.Time{}),
		field.String("repeat").Optional().GoType(null.String{}),
		field.String("subtasks").GoType(&types.Subtasks{}),
		field.Bool("is_completed").Default(false),
		field.Bool("is_deleted").Default(false),
		field.Time("last_modified").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "task"},
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("list", List.Type).
			Ref("tasks").
			Unique().
			Required().Field("list_id"),
	}
}
