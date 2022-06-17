// Code generated by entc, DO NOT EDIT.

package task

import (
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldListID holds the string denoting the list_id field in the database.
	FieldListID = "list_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldReminder holds the string denoting the reminder field in the database.
	FieldReminder = "reminder"
	// FieldRepeat holds the string denoting the repeat field in the database.
	FieldRepeat = "repeat"
	// FieldSubtasks holds the string denoting the subtasks field in the database.
	FieldSubtasks = "subtasks"
	// FieldIsCompleted holds the string denoting the is_completed field in the database.
	FieldIsCompleted = "is_completed"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldLastModified holds the string denoting the last_modified field in the database.
	FieldLastModified = "last_modified"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeList holds the string denoting the list edge name in mutations.
	EdgeList = "list"
	// Table holds the table name of the task in the database.
	Table = "task"
	// ListTable is the table that holds the list relation/edge.
	ListTable = "task"
	// ListInverseTable is the table name for the List entity.
	// It exists in this package in order to avoid circular dependency with the "list" package.
	ListInverseTable = "list"
	// ListColumn is the table column denoting the list relation/edge.
	ListColumn = "list_id"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldListID,
	FieldDescription,
	FieldReminder,
	FieldRepeat,
	FieldSubtasks,
	FieldIsCompleted,
	FieldIsDeleted,
	FieldLastModified,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsCompleted holds the default value on creation for the "is_completed" field.
	DefaultIsCompleted bool
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
	// DefaultLastModified holds the default value on creation for the "last_modified" field.
	DefaultLastModified func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
