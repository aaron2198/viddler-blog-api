// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/ent/article"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/ent/useroptions"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VideoURL holds the value of the "video_url" field.
	VideoURL string `json:"video_url,omitempty"`
	// VideoID holds the value of the "video_id" field.
	VideoID string `json:"video_id,omitempty"`
	// Uploader holds the value of the "uploader" field.
	Uploader string `json:"uploader,omitempty"`
	// UploaderURL holds the value of the "uploader_url" field.
	UploaderURL string `json:"uploader_url,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Thumbnail holds the value of the "thumbnail" field.
	Thumbnail string `json:"thumbnail,omitempty"`
	// HTML holds the value of the "html" field.
	HTML string `json:"html,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleQuery when eager-loading is set.
	Edges                ArticleEdges `json:"edges"`
	user_options_article *int
	selectValues         sql.SelectValues
}

// ArticleEdges holds the relations/edges for other nodes in the graph.
type ArticleEdges struct {
	// UserOptions holds the value of the user_options edge.
	UserOptions *UserOptions `json:"user_options,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOptionsOrErr returns the UserOptions value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleEdges) UserOptionsOrErr() (*UserOptions, error) {
	if e.UserOptions != nil {
		return e.UserOptions, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: useroptions.Label}
	}
	return nil, &NotLoadedError{edge: "user_options"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			values[i] = new(sql.NullInt64)
		case article.FieldVideoURL, article.FieldVideoID, article.FieldUploader, article.FieldUploaderURL, article.FieldDescription, article.FieldTitle, article.FieldThumbnail, article.FieldHTML:
			values[i] = new(sql.NullString)
		case article.ForeignKeys[0]: // user_options_article
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case article.FieldVideoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_url", values[i])
			} else if value.Valid {
				a.VideoURL = value.String
			}
		case article.FieldVideoID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				a.VideoID = value.String
			}
		case article.FieldUploader:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uploader", values[i])
			} else if value.Valid {
				a.Uploader = value.String
			}
		case article.FieldUploaderURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uploader_url", values[i])
			} else if value.Valid {
				a.UploaderURL = value.String
			}
		case article.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldThumbnail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail", values[i])
			} else if value.Valid {
				a.Thumbnail = value.String
			}
		case article.FieldHTML:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html", values[i])
			} else if value.Valid {
				a.HTML = value.String
			}
		case article.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_options_article", value)
			} else if value.Valid {
				a.user_options_article = new(int)
				*a.user_options_article = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Article.
// This includes values selected through modifiers, order, etc.
func (a *Article) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUserOptions queries the "user_options" edge of the Article entity.
func (a *Article) QueryUserOptions() *UserOptionsQuery {
	return NewArticleClient(a.config).QueryUserOptions(a)
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return NewArticleClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("video_url=")
	builder.WriteString(a.VideoURL)
	builder.WriteString(", ")
	builder.WriteString("video_id=")
	builder.WriteString(a.VideoID)
	builder.WriteString(", ")
	builder.WriteString("uploader=")
	builder.WriteString(a.Uploader)
	builder.WriteString(", ")
	builder.WriteString("uploader_url=")
	builder.WriteString(a.UploaderURL)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("thumbnail=")
	builder.WriteString(a.Thumbnail)
	builder.WriteString(", ")
	builder.WriteString("html=")
	builder.WriteString(a.HTML)
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article
