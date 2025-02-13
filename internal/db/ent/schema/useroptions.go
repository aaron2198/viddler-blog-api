package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserOptions holds the schema definition for the UserOptions entity.
type UserOptions struct {
	ent.Schema
}

// Fields of the UserOptions.
func (UserOptions) Fields() []ent.Field {
	return []ent.Field{
		field.String("video_url").NotEmpty(),
		field.String("user_style_prompt").Optional(),
		field.String("client").Optional(),
		field.String("model").Optional(),
		field.String("mode").Optional(),
		field.Bool("chapters_as_sections").Default(false),
		field.Bool("embed_video").Default(false),
		field.Bool("include_description").Default(false),
		field.Bool("include_tags").Default(false),
	}
}

// Edges of the UserOptions.
func (UserOptions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("article", Article.Type).Unique(),
		edge.To("phase_options", PhaseOptions.Type),
	}
}
