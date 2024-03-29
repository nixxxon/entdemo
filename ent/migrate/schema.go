// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "other_id", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
	}
	// TodoHistoryColumns holds the columns for the "todo_history" table.
	TodoHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeUUID, Nullable: true},
		{Name: "other_id", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString},
	}
	// TodoHistoryTable holds the schema information for the "todo_history" table.
	TodoHistoryTable = &schema.Table{
		Name:       "todo_history",
		Columns:    TodoHistoryColumns,
		PrimaryKey: []*schema.Column{TodoHistoryColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
		TodoHistoryTable,
	}
)

func init() {
	TodoHistoryTable.Annotation = &entsql.Annotation{
		Table: "todo_history",
	}
}
