// Code generated by entc, DO NOT EDIT.

package team

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTeamNumber holds the string denoting the team_number field in the database.
	FieldTeamNumber = "team_number"
	// FieldVars holds the string denoting the vars field in the database.
	FieldVars = "vars"
	// EdgeTeamToBuild holds the string denoting the teamtobuild edge name in mutations.
	EdgeTeamToBuild = "TeamToBuild"
	// EdgeTeamToStatus holds the string denoting the teamtostatus edge name in mutations.
	EdgeTeamToStatus = "TeamToStatus"
	// EdgeTeamToProvisionedNetwork holds the string denoting the teamtoprovisionednetwork edge name in mutations.
	EdgeTeamToProvisionedNetwork = "TeamToProvisionedNetwork"
	// EdgeTeamToPlan holds the string denoting the teamtoplan edge name in mutations.
	EdgeTeamToPlan = "TeamToPlan"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// TeamToBuildTable is the table that holds the TeamToBuild relation/edge.
	TeamToBuildTable = "teams"
	// TeamToBuildInverseTable is the table name for the Build entity.
	// It exists in this package in order to avoid circular dependency with the "build" package.
	TeamToBuildInverseTable = "builds"
	// TeamToBuildColumn is the table column denoting the TeamToBuild relation/edge.
	TeamToBuildColumn = "team_team_to_build"
	// TeamToStatusTable is the table that holds the TeamToStatus relation/edge.
	TeamToStatusTable = "status"
	// TeamToStatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	TeamToStatusInverseTable = "status"
	// TeamToStatusColumn is the table column denoting the TeamToStatus relation/edge.
	TeamToStatusColumn = "team_team_to_status"
	// TeamToProvisionedNetworkTable is the table that holds the TeamToProvisionedNetwork relation/edge.
	TeamToProvisionedNetworkTable = "provisioned_networks"
	// TeamToProvisionedNetworkInverseTable is the table name for the ProvisionedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "provisionednetwork" package.
	TeamToProvisionedNetworkInverseTable = "provisioned_networks"
	// TeamToProvisionedNetworkColumn is the table column denoting the TeamToProvisionedNetwork relation/edge.
	TeamToProvisionedNetworkColumn = "provisioned_network_provisioned_network_to_team"
	// TeamToPlanTable is the table that holds the TeamToPlan relation/edge.
	TeamToPlanTable = "teams"
	// TeamToPlanInverseTable is the table name for the Plan entity.
	// It exists in this package in order to avoid circular dependency with the "plan" package.
	TeamToPlanInverseTable = "plans"
	// TeamToPlanColumn is the table column denoting the TeamToPlan relation/edge.
	TeamToPlanColumn = "plan_plan_to_team"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldTeamNumber,
	FieldVars,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "teams"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"plan_plan_to_team",
	"team_team_to_build",
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