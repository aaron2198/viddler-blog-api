package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("video_url").NotEmpty(),
		field.String("video_id").NotEmpty(),
		field.String("uploader").NotEmpty(),
		field.String("uploader_url").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("title").NotEmpty(),
		field.String("thumbnail").NotEmpty(),
		field.String("html"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_options", UserOptions.Type).Ref("article").Unique(),
	}
}
