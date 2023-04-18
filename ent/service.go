// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"services/ent/service"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Service is the model entity for the Service schema.
type Service struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedOn holds the value of the "created_on" field.
	CreatedOn time.Time `json:"created_on,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServiceQuery when eager-loading is set.
	Edges        ServiceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ServiceEdges holds the relations/edges for other nodes in the graph.
type ServiceEdges struct {
	// Versions holds the value of the versions edge.
	Versions []*ServiceVersion `json:"versions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VersionsOrErr returns the Versions value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceEdges) VersionsOrErr() ([]*ServiceVersion, error) {
	if e.loadedTypes[0] {
		return e.Versions, nil
	}
	return nil, &NotLoadedError{edge: "versions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Service) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			values[i] = new(sql.NullInt64)
		case service.FieldName, service.FieldDescription:
			values[i] = new(sql.NullString)
		case service.FieldCreatedOn:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Service fields.
func (s *Service) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case service.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case service.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case service.FieldCreatedOn:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_on", values[i])
			} else if value.Valid {
				s.CreatedOn = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Service.
// This includes values selected through modifiers, order, etc.
func (s *Service) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryVersions queries the "versions" edge of the Service entity.
func (s *Service) QueryVersions() *ServiceVersionQuery {
	return NewServiceClient(s.config).QueryVersions(s)
}

// Update returns a builder for updating this Service.
// Note that you need to call Service.Unwrap() before calling this method if this Service
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Service) Update() *ServiceUpdateOne {
	return NewServiceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Service entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Service) Unwrap() *Service {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Service is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Service) String() string {
	var builder strings.Builder
	builder.WriteString("Service(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("created_on=")
	builder.WriteString(s.CreatedOn.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Services is a parsable slice of Service.
type Services []*Service
