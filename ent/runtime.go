// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/nixxxon/entdemo/ent/schema"
	"github.com/nixxxon/entdemo/ent/todo"
	"github.com/nixxxon/entdemo/ent/todohack"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescName is the schema descriptor for name field.
	todoDescName := todoFields[2].Descriptor()
	// todo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	todo.NameValidator = todoDescName.Validators[0].(func(string) error)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoFields[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() uuid.UUID)
	todohackFields := schema.TodoHack{}.Fields()
	_ = todohackFields
	// todohackDescHistoryTime is the schema descriptor for history_time field.
	todohackDescHistoryTime := todohackFields[0].Descriptor()
	// todohack.DefaultHistoryTime holds the default value on creation for the history_time field.
	todohack.DefaultHistoryTime = todohackDescHistoryTime.Default.(func() time.Time)
	// todohackDescName is the schema descriptor for name field.
	todohackDescName := todohackFields[5].Descriptor()
	// todohack.NameValidator is a validator for the "name" field. It is called by the builders before save.
	todohack.NameValidator = todohackDescName.Validators[0].(func(string) error)
	// todohackDescID is the schema descriptor for id field.
	todohackDescID := todohackFields[3].Descriptor()
	// todohack.DefaultID holds the default value on creation for the id field.
	todohack.DefaultID = todohackDescID.Default.(func() uuid.UUID)
}