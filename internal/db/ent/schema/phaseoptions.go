package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PhaseOptions holds the schema definition for the PhaseOptions entity.
type PhaseOptions struct {
	ent.Schema
}

// Fields of the PhaseOptions.
func (PhaseOptions) Fields() []ent.Field {
	return []ent.Field{
		field.String("phase_name").NotEmpty(),
		field.String("client").Optional(),
		field.String("model").Optional(),
		field.String("prompt").Optional(),
	}
}

// Edges of the PhaseOptions.
func (PhaseOptions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_options", UserOptions.Type).Ref("phase_options").Unique(),
	}
}
