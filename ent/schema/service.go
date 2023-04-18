package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name"),
		field.String("description"),
		field.Time("created_on"),
		// field.Int("versions").Positive(),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("versions", ServiceVersion.Type),
	}
}

func (Service) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Services"},
	}
}
