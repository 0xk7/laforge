// Code generated by entc, DO NOT EDIT.

package fileextract

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the fileextract type in the database.
	Label = "file_extract"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHclID holds the string denoting the hcl_id field in the database.
	FieldHclID = "hcl_id"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldDestination holds the string denoting the destination field in the database.
	FieldDestination = "destination"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// EdgeFileExtractToEnvironment holds the string denoting the fileextracttoenvironment edge name in mutations.
	EdgeFileExtractToEnvironment = "FileExtractToEnvironment"
	// Table holds the table name of the fileextract in the database.
	Table = "file_extracts"
	// FileExtractToEnvironmentTable is the table that holds the FileExtractToEnvironment relation/edge.
	FileExtractToEnvironmentTable = "file_extracts"
	// FileExtractToEnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	FileExtractToEnvironmentInverseTable = "environments"
	// FileExtractToEnvironmentColumn is the table column denoting the FileExtractToEnvironment relation/edge.
	FileExtractToEnvironmentColumn = "environment_environment_to_file_extract"
)

// Columns holds all SQL columns for fileextract fields.
var Columns = []string{
	FieldID,
	FieldHclID,
	FieldSource,
	FieldDestination,
	FieldType,
	FieldTags,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "file_extracts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"environment_environment_to_file_extract",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)