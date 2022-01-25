package schema

import (
	"entgo.io/ent"
	"github.com/maguowei/example/internal/pkg/mixin"
)

// Feedback holds the schema definition for the Feedback entity.
type Feedback struct {
	ent.Schema
}

func (Feedback) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Feedback.
func (Feedback) Fields() []ent.Field {
	return []ent.Field{
	}
}

// Edges of the Feedback.
func (Feedback) Edges() []ent.Edge {
	return nil
}
