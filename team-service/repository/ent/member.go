// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"team-service/repository/ent/member"
	"team-service/repository/ent/team"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"memberId"`
	// TeamID holds the value of the "team_id" field.
	TeamID *string `json:"teamId"`
	// AssignedBy holds the value of the "assigned_by" field.
	AssignedBy *string `json:"assignedBy"`
	// ApprovedBy holds the value of the "approved_by" field.
	ApprovedBy *string `json:"approvedBy"`
	// Status holds the value of the "status" field.
	Status member.Status `json:"status"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberQuery when eager-loading is set.
	Edges        MemberEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MemberEdges holds the relations/edges for other nodes in the graph.
type MemberEdges struct {
	// Team holds the value of the team edge.
	Team []*Team `json:"team,omitempty"`
	// Teams holds the value of the teams edge.
	Teams *Team `json:"teams,omitempty"`
	// Assigned holds the value of the assigned edge.
	Assigned *Member `json:"assigned,omitempty"`
	// Member holds the value of the member edge.
	Member []*Member `json:"member,omitempty"`
	// Approved holds the value of the approved edge.
	Approved *Member `json:"approved,omitempty"`
	// Approve holds the value of the approve edge.
	Approve []*Member `json:"approve,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) TeamOrErr() ([]*Team, error) {
	if e.loadedTypes[0] {
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberEdges) TeamsOrErr() (*Team, error) {
	if e.loadedTypes[1] {
		if e.Teams == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
}

// AssignedOrErr returns the Assigned value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberEdges) AssignedOrErr() (*Member, error) {
	if e.loadedTypes[2] {
		if e.Assigned == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Assigned, nil
	}
	return nil, &NotLoadedError{edge: "assigned"}
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) MemberOrErr() ([]*Member, error) {
	if e.loadedTypes[3] {
		return e.Member, nil
	}
	return nil, &NotLoadedError{edge: "member"}
}

// ApprovedOrErr returns the Approved value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberEdges) ApprovedOrErr() (*Member, error) {
	if e.loadedTypes[4] {
		if e.Approved == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Approved, nil
	}
	return nil, &NotLoadedError{edge: "approved"}
}

// ApproveOrErr returns the Approve value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) ApproveOrErr() ([]*Member, error) {
	if e.loadedTypes[5] {
		return e.Approve, nil
	}
	return nil, &NotLoadedError{edge: "approve"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case member.FieldID, member.FieldTeamID, member.FieldAssignedBy, member.FieldApprovedBy, member.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case member.FieldTeamID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				m.TeamID = new(string)
				*m.TeamID = value.String
			}
		case member.FieldAssignedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field assigned_by", values[i])
			} else if value.Valid {
				m.AssignedBy = new(string)
				*m.AssignedBy = value.String
			}
		case member.FieldApprovedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field approved_by", values[i])
			} else if value.Valid {
				m.ApprovedBy = new(string)
				*m.ApprovedBy = value.String
			}
		case member.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = member.Status(value.String)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Member.
// This includes values selected through modifiers, order, etc.
func (m *Member) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryTeam queries the "team" edge of the Member entity.
func (m *Member) QueryTeam() *TeamQuery {
	return NewMemberClient(m.config).QueryTeam(m)
}

// QueryTeams queries the "teams" edge of the Member entity.
func (m *Member) QueryTeams() *TeamQuery {
	return NewMemberClient(m.config).QueryTeams(m)
}

// QueryAssigned queries the "assigned" edge of the Member entity.
func (m *Member) QueryAssigned() *MemberQuery {
	return NewMemberClient(m.config).QueryAssigned(m)
}

// QueryMember queries the "member" edge of the Member entity.
func (m *Member) QueryMember() *MemberQuery {
	return NewMemberClient(m.config).QueryMember(m)
}

// QueryApproved queries the "approved" edge of the Member entity.
func (m *Member) QueryApproved() *MemberQuery {
	return NewMemberClient(m.config).QueryApproved(m)
}

// QueryApprove queries the "approve" edge of the Member entity.
func (m *Member) QueryApprove() *MemberQuery {
	return NewMemberClient(m.config).QueryApprove(m)
}

// Update returns a builder for updating this Member.
// Note that you need to call Member.Unwrap() before calling this method if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return NewMemberClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Member entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	if v := m.TeamID; v != nil {
		builder.WriteString("team_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.AssignedBy; v != nil {
		builder.WriteString("assigned_by=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.ApprovedBy; v != nil {
		builder.WriteString("approved_by=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member
