// Code generated by ent, DO NOT EDIT.

package ent

import (
	"services/ent/schema"
	"services/ent/service"
	"services/ent/serviceversion"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	serviceFields := schema.Service{}.Fields()
	_ = serviceFields
	// serviceDescID is the schema descriptor for id field.
	serviceDescID := serviceFields[0].Descriptor()
	// service.IDValidator is a validator for the "id" field. It is called by the builders before save.
	service.IDValidator = serviceDescID.Validators[0].(func(int) error)
	serviceversionFields := schema.ServiceVersion{}.Fields()
	_ = serviceversionFields
	// serviceversionDescID is the schema descriptor for id field.
	serviceversionDescID := serviceversionFields[0].Descriptor()
	// serviceversion.IDValidator is a validator for the "id" field. It is called by the builders before save.
	serviceversion.IDValidator = serviceversionDescID.Validators[0].(func(int) error)
}
