package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ServiceVersion holds the schema definition for the ServiceVersion entity.
type ServiceVersion struct {
	ent.Schema
}

// Fields of the ServiceVersion.
func (ServiceVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name"),
		field.Time("created_on"),
		field.Int("service_id"),
	}
}

// Edges of the ServiceVersion.
func (ServiceVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("service", Service.Type).
			Ref("versions").
			Field("service_id").
			Unique().Required(),
	}
}

func (ServiceVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ServiceVersions"},
	}
}
