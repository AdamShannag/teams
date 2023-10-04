// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"project-service/repository/ent/attachment"
	"project-service/repository/ent/member"
	"project-service/repository/ent/project"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttachmentCreate is the builder for creating a Attachment entity.
type AttachmentCreate struct {
	config
	mutation *AttachmentMutation
	hooks    []Hook
}

// SetAttachment sets the "attachment" field.
func (ac *AttachmentCreate) SetAttachment(s string) *AttachmentCreate {
	ac.mutation.SetAttachment(s)
	return ac
}

// SetProjectID sets the "project_id" field.
func (ac *AttachmentCreate) SetProjectID(s string) *AttachmentCreate {
	ac.mutation.SetProjectID(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AttachmentCreate) SetDescription(s string) *AttachmentCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableDescription(s *string) *AttachmentCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetAddedBy sets the "added_by" field.
func (ac *AttachmentCreate) SetAddedBy(s string) *AttachmentCreate {
	ac.mutation.SetAddedBy(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AttachmentCreate) SetCreatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCreatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AttachmentCreate) SetUpdatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableUpdatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AttachmentCreate) SetID(s string) *AttachmentCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetInsertionID sets the "insertion" edge to the Member entity by ID.
func (ac *AttachmentCreate) SetInsertionID(id string) *AttachmentCreate {
	ac.mutation.SetInsertionID(id)
	return ac
}

// SetInsertion sets the "insertion" edge to the Member entity.
func (ac *AttachmentCreate) SetInsertion(m *Member) *AttachmentCreate {
	return ac.SetInsertionID(m.ID)
}

// SetAssignationID sets the "assignation" edge to the Project entity by ID.
func (ac *AttachmentCreate) SetAssignationID(id string) *AttachmentCreate {
	ac.mutation.SetAssignationID(id)
	return ac
}

// SetAssignation sets the "assignation" edge to the Project entity.
func (ac *AttachmentCreate) SetAssignation(p *Project) *AttachmentCreate {
	return ac.SetAssignationID(p.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (ac *AttachmentCreate) Mutation() *AttachmentMutation {
	return ac.mutation
}

// Save creates the Attachment in the database.
func (ac *AttachmentCreate) Save(ctx context.Context) (*Attachment, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttachmentCreate) SaveX(ctx context.Context) *Attachment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttachmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttachmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttachmentCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := attachment.DefaultCreatedAt
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := attachment.DefaultUpdatedAt
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttachmentCreate) check() error {
	if _, ok := ac.mutation.Attachment(); !ok {
		return &ValidationError{Name: "attachment", err: errors.New(`ent: missing required field "Attachment.attachment"`)}
	}
	if _, ok := ac.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "Attachment.project_id"`)}
	}
	if _, ok := ac.mutation.AddedBy(); !ok {
		return &ValidationError{Name: "added_by", err: errors.New(`ent: missing required field "Attachment.added_by"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Attachment.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Attachment.updated_at"`)}
	}
	if _, ok := ac.mutation.InsertionID(); !ok {
		return &ValidationError{Name: "insertion", err: errors.New(`ent: missing required edge "Attachment.insertion"`)}
	}
	if _, ok := ac.mutation.AssignationID(); !ok {
		return &ValidationError{Name: "assignation", err: errors.New(`ent: missing required edge "Attachment.assignation"`)}
	}
	return nil
}

func (ac *AttachmentCreate) sqlSave(ctx context.Context) (*Attachment, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Attachment.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AttachmentCreate) createSpec() (*Attachment, *sqlgraph.CreateSpec) {
	var (
		_node = &Attachment{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(attachment.Table, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeString))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Attachment(); ok {
		_spec.SetField(attachment.FieldAttachment, field.TypeString, value)
		_node.Attachment = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(attachment.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.InsertionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attachment.InsertionTable,
			Columns: []string{attachment.InsertionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AddedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AssignationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attachment.AssignationTable,
			Columns: []string{attachment.AssignationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttachmentCreateBulk is the builder for creating many Attachment entities in bulk.
type AttachmentCreateBulk struct {
	config
	err      error
	builders []*AttachmentCreate
}

// Save creates the Attachment entities in the database.
func (acb *AttachmentCreateBulk) Save(ctx context.Context) ([]*Attachment, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attachment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttachmentCreateBulk) SaveX(ctx context.Context) []*Attachment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttachmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttachmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}