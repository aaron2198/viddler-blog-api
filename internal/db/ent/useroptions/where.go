// Code generated by ent, DO NOT EDIT.

package useroptions

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLTE(FieldID, id))
}

// VideoURL applies equality check predicate on the "video_url" field. It's identical to VideoURLEQ.
func VideoURL(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldVideoURL, v))
}

// UserStylePrompt applies equality check predicate on the "user_style_prompt" field. It's identical to UserStylePromptEQ.
func UserStylePrompt(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldUserStylePrompt, v))
}

// Client applies equality check predicate on the "client" field. It's identical to ClientEQ.
func Client(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldClient, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldModel, v))
}

// Mode applies equality check predicate on the "mode" field. It's identical to ModeEQ.
func Mode(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldMode, v))
}

// ChaptersAsSections applies equality check predicate on the "chapters_as_sections" field. It's identical to ChaptersAsSectionsEQ.
func ChaptersAsSections(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldChaptersAsSections, v))
}

// EmbedVideo applies equality check predicate on the "embed_video" field. It's identical to EmbedVideoEQ.
func EmbedVideo(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldEmbedVideo, v))
}

// IncludeDescription applies equality check predicate on the "include_description" field. It's identical to IncludeDescriptionEQ.
func IncludeDescription(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldIncludeDescription, v))
}

// IncludeTags applies equality check predicate on the "include_tags" field. It's identical to IncludeTagsEQ.
func IncludeTags(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldIncludeTags, v))
}

// VideoURLEQ applies the EQ predicate on the "video_url" field.
func VideoURLEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldVideoURL, v))
}

// VideoURLNEQ applies the NEQ predicate on the "video_url" field.
func VideoURLNEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldVideoURL, v))
}

// VideoURLIn applies the In predicate on the "video_url" field.
func VideoURLIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIn(FieldVideoURL, vs...))
}

// VideoURLNotIn applies the NotIn predicate on the "video_url" field.
func VideoURLNotIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotIn(FieldVideoURL, vs...))
}

// VideoURLGT applies the GT predicate on the "video_url" field.
func VideoURLGT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGT(FieldVideoURL, v))
}

// VideoURLGTE applies the GTE predicate on the "video_url" field.
func VideoURLGTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGTE(FieldVideoURL, v))
}

// VideoURLLT applies the LT predicate on the "video_url" field.
func VideoURLLT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLT(FieldVideoURL, v))
}

// VideoURLLTE applies the LTE predicate on the "video_url" field.
func VideoURLLTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLTE(FieldVideoURL, v))
}

// VideoURLContains applies the Contains predicate on the "video_url" field.
func VideoURLContains(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContains(FieldVideoURL, v))
}

// VideoURLHasPrefix applies the HasPrefix predicate on the "video_url" field.
func VideoURLHasPrefix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasPrefix(FieldVideoURL, v))
}

// VideoURLHasSuffix applies the HasSuffix predicate on the "video_url" field.
func VideoURLHasSuffix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasSuffix(FieldVideoURL, v))
}

// VideoURLEqualFold applies the EqualFold predicate on the "video_url" field.
func VideoURLEqualFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEqualFold(FieldVideoURL, v))
}

// VideoURLContainsFold applies the ContainsFold predicate on the "video_url" field.
func VideoURLContainsFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContainsFold(FieldVideoURL, v))
}

// UserStylePromptEQ applies the EQ predicate on the "user_style_prompt" field.
func UserStylePromptEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldUserStylePrompt, v))
}

// UserStylePromptNEQ applies the NEQ predicate on the "user_style_prompt" field.
func UserStylePromptNEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldUserStylePrompt, v))
}

// UserStylePromptIn applies the In predicate on the "user_style_prompt" field.
func UserStylePromptIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIn(FieldUserStylePrompt, vs...))
}

// UserStylePromptNotIn applies the NotIn predicate on the "user_style_prompt" field.
func UserStylePromptNotIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotIn(FieldUserStylePrompt, vs...))
}

// UserStylePromptGT applies the GT predicate on the "user_style_prompt" field.
func UserStylePromptGT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGT(FieldUserStylePrompt, v))
}

// UserStylePromptGTE applies the GTE predicate on the "user_style_prompt" field.
func UserStylePromptGTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGTE(FieldUserStylePrompt, v))
}

// UserStylePromptLT applies the LT predicate on the "user_style_prompt" field.
func UserStylePromptLT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLT(FieldUserStylePrompt, v))
}

// UserStylePromptLTE applies the LTE predicate on the "user_style_prompt" field.
func UserStylePromptLTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLTE(FieldUserStylePrompt, v))
}

// UserStylePromptContains applies the Contains predicate on the "user_style_prompt" field.
func UserStylePromptContains(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContains(FieldUserStylePrompt, v))
}

// UserStylePromptHasPrefix applies the HasPrefix predicate on the "user_style_prompt" field.
func UserStylePromptHasPrefix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasPrefix(FieldUserStylePrompt, v))
}

// UserStylePromptHasSuffix applies the HasSuffix predicate on the "user_style_prompt" field.
func UserStylePromptHasSuffix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasSuffix(FieldUserStylePrompt, v))
}

// UserStylePromptIsNil applies the IsNil predicate on the "user_style_prompt" field.
func UserStylePromptIsNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIsNull(FieldUserStylePrompt))
}

// UserStylePromptNotNil applies the NotNil predicate on the "user_style_prompt" field.
func UserStylePromptNotNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotNull(FieldUserStylePrompt))
}

// UserStylePromptEqualFold applies the EqualFold predicate on the "user_style_prompt" field.
func UserStylePromptEqualFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEqualFold(FieldUserStylePrompt, v))
}

// UserStylePromptContainsFold applies the ContainsFold predicate on the "user_style_prompt" field.
func UserStylePromptContainsFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContainsFold(FieldUserStylePrompt, v))
}

// ClientEQ applies the EQ predicate on the "client" field.
func ClientEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldClient, v))
}

// ClientNEQ applies the NEQ predicate on the "client" field.
func ClientNEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldClient, v))
}

// ClientIn applies the In predicate on the "client" field.
func ClientIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIn(FieldClient, vs...))
}

// ClientNotIn applies the NotIn predicate on the "client" field.
func ClientNotIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotIn(FieldClient, vs...))
}

// ClientGT applies the GT predicate on the "client" field.
func ClientGT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGT(FieldClient, v))
}

// ClientGTE applies the GTE predicate on the "client" field.
func ClientGTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGTE(FieldClient, v))
}

// ClientLT applies the LT predicate on the "client" field.
func ClientLT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLT(FieldClient, v))
}

// ClientLTE applies the LTE predicate on the "client" field.
func ClientLTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLTE(FieldClient, v))
}

// ClientContains applies the Contains predicate on the "client" field.
func ClientContains(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContains(FieldClient, v))
}

// ClientHasPrefix applies the HasPrefix predicate on the "client" field.
func ClientHasPrefix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasPrefix(FieldClient, v))
}

// ClientHasSuffix applies the HasSuffix predicate on the "client" field.
func ClientHasSuffix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasSuffix(FieldClient, v))
}

// ClientIsNil applies the IsNil predicate on the "client" field.
func ClientIsNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIsNull(FieldClient))
}

// ClientNotNil applies the NotNil predicate on the "client" field.
func ClientNotNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotNull(FieldClient))
}

// ClientEqualFold applies the EqualFold predicate on the "client" field.
func ClientEqualFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEqualFold(FieldClient, v))
}

// ClientContainsFold applies the ContainsFold predicate on the "client" field.
func ClientContainsFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContainsFold(FieldClient, v))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasSuffix(FieldModel, v))
}

// ModelIsNil applies the IsNil predicate on the "model" field.
func ModelIsNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIsNull(FieldModel))
}

// ModelNotNil applies the NotNil predicate on the "model" field.
func ModelNotNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotNull(FieldModel))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContainsFold(FieldModel, v))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldMode, v))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldMode, v))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIn(FieldMode, vs...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotIn(FieldMode, vs...))
}

// ModeGT applies the GT predicate on the "mode" field.
func ModeGT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGT(FieldMode, v))
}

// ModeGTE applies the GTE predicate on the "mode" field.
func ModeGTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldGTE(FieldMode, v))
}

// ModeLT applies the LT predicate on the "mode" field.
func ModeLT(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLT(FieldMode, v))
}

// ModeLTE applies the LTE predicate on the "mode" field.
func ModeLTE(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldLTE(FieldMode, v))
}

// ModeContains applies the Contains predicate on the "mode" field.
func ModeContains(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContains(FieldMode, v))
}

// ModeHasPrefix applies the HasPrefix predicate on the "mode" field.
func ModeHasPrefix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasPrefix(FieldMode, v))
}

// ModeHasSuffix applies the HasSuffix predicate on the "mode" field.
func ModeHasSuffix(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldHasSuffix(FieldMode, v))
}

// ModeIsNil applies the IsNil predicate on the "mode" field.
func ModeIsNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldIsNull(FieldMode))
}

// ModeNotNil applies the NotNil predicate on the "mode" field.
func ModeNotNil() predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNotNull(FieldMode))
}

// ModeEqualFold applies the EqualFold predicate on the "mode" field.
func ModeEqualFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEqualFold(FieldMode, v))
}

// ModeContainsFold applies the ContainsFold predicate on the "mode" field.
func ModeContainsFold(v string) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldContainsFold(FieldMode, v))
}

// ChaptersAsSectionsEQ applies the EQ predicate on the "chapters_as_sections" field.
func ChaptersAsSectionsEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldChaptersAsSections, v))
}

// ChaptersAsSectionsNEQ applies the NEQ predicate on the "chapters_as_sections" field.
func ChaptersAsSectionsNEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldChaptersAsSections, v))
}

// EmbedVideoEQ applies the EQ predicate on the "embed_video" field.
func EmbedVideoEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldEmbedVideo, v))
}

// EmbedVideoNEQ applies the NEQ predicate on the "embed_video" field.
func EmbedVideoNEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldEmbedVideo, v))
}

// IncludeDescriptionEQ applies the EQ predicate on the "include_description" field.
func IncludeDescriptionEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldIncludeDescription, v))
}

// IncludeDescriptionNEQ applies the NEQ predicate on the "include_description" field.
func IncludeDescriptionNEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldIncludeDescription, v))
}

// IncludeTagsEQ applies the EQ predicate on the "include_tags" field.
func IncludeTagsEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldEQ(FieldIncludeTags, v))
}

// IncludeTagsNEQ applies the NEQ predicate on the "include_tags" field.
func IncludeTagsNEQ(v bool) predicate.UserOptions {
	return predicate.UserOptions(sql.FieldNEQ(FieldIncludeTags, v))
}

// HasArticle applies the HasEdge predicate on the "article" edge.
func HasArticle() predicate.UserOptions {
	return predicate.UserOptions(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ArticleTable, ArticleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticleWith applies the HasEdge predicate on the "article" edge with a given conditions (other predicates).
func HasArticleWith(preds ...predicate.Article) predicate.UserOptions {
	return predicate.UserOptions(func(s *sql.Selector) {
		step := newArticleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPhaseOptions applies the HasEdge predicate on the "phase_options" edge.
func HasPhaseOptions() predicate.UserOptions {
	return predicate.UserOptions(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PhaseOptionsTable, PhaseOptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPhaseOptionsWith applies the HasEdge predicate on the "phase_options" edge with a given conditions (other predicates).
func HasPhaseOptionsWith(preds ...predicate.PhaseOptions) predicate.UserOptions {
	return predicate.UserOptions(func(s *sql.Selector) {
		step := newPhaseOptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserOptions) predicate.UserOptions {
	return predicate.UserOptions(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserOptions) predicate.UserOptions {
	return predicate.UserOptions(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserOptions) predicate.UserOptions {
	return predicate.UserOptions(sql.NotPredicates(p))
}
