// Code generated by ent, DO NOT EDIT.

package member

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeTeamID holds the string denoting the team_id edge name in mutations.
	EdgeTeamID = "team_id"
	// Table holds the table name of the member in the database.
	Table = "members"
	// TeamIDTable is the table that holds the team_id relation/edge.
	TeamIDTable = "members"
	// TeamIDInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamIDInverseTable = "teams"
	// TeamIDColumn is the table column denoting the team_id relation/edge.
	TeamIDColumn = "team_members"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "members"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"team_members",
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

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusFREE    Status = "FREE"
	StatusPENDING Status = "PENDING"
	StatusIN_TEAM Status = "IN_TEAM"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusFREE, StatusPENDING, StatusIN_TEAM:
		return nil
	default:
		return fmt.Errorf("member: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Member queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTeamIDField orders the results by team_id field.
func ByTeamIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamIDStep(), sql.OrderByField(field, opts...))
	}
}
func newTeamIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TeamIDTable, TeamIDColumn),
	)
}
