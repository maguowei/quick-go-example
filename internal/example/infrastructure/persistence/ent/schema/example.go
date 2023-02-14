package schema

import (
	"entgo.io/ent"
	"example/internal/pkg/mixin"
)

// Example holds the schema definition for the Example entity.
type Example struct {
	ent.Schema
}

func (Example) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Example.
func (Example) Fields() []ent.Field {
	return []ent.Field{
	}
}

// Edges of the Example.
func (Example) Edges() []ent.Edge {
	return nil
}
