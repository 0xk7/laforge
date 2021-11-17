// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/identity"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// IdentityUpdate is the builder for updating Identity entities.
type IdentityUpdate struct {
	config
	hooks    []Hook
	mutation *IdentityMutation
}

// Where appends a list predicates to the IdentityUpdate builder.
func (iu *IdentityUpdate) Where(ps ...predicate.Identity) *IdentityUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetHclID sets the "hcl_id" field.
func (iu *IdentityUpdate) SetHclID(s string) *IdentityUpdate {
	iu.mutation.SetHclID(s)
	return iu
}

// SetFirstName sets the "first_name" field.
func (iu *IdentityUpdate) SetFirstName(s string) *IdentityUpdate {
	iu.mutation.SetFirstName(s)
	return iu
}

// SetLastName sets the "last_name" field.
func (iu *IdentityUpdate) SetLastName(s string) *IdentityUpdate {
	iu.mutation.SetLastName(s)
	return iu
}

// SetEmail sets the "email" field.
func (iu *IdentityUpdate) SetEmail(s string) *IdentityUpdate {
	iu.mutation.SetEmail(s)
	return iu
}

// SetPassword sets the "password" field.
func (iu *IdentityUpdate) SetPassword(s string) *IdentityUpdate {
	iu.mutation.SetPassword(s)
	return iu
}

// SetDescription sets the "description" field.
func (iu *IdentityUpdate) SetDescription(s string) *IdentityUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetAvatarFile sets the "avatar_file" field.
func (iu *IdentityUpdate) SetAvatarFile(s string) *IdentityUpdate {
	iu.mutation.SetAvatarFile(s)
	return iu
}

// SetVars sets the "vars" field.
func (iu *IdentityUpdate) SetVars(m map[string]string) *IdentityUpdate {
	iu.mutation.SetVars(m)
	return iu
}

// SetTags sets the "tags" field.
func (iu *IdentityUpdate) SetTags(m map[string]string) *IdentityUpdate {
	iu.mutation.SetTags(m)
	return iu
}

// SetIdentityToEnvironmentID sets the "IdentityToEnvironment" edge to the Environment entity by ID.
func (iu *IdentityUpdate) SetIdentityToEnvironmentID(id uuid.UUID) *IdentityUpdate {
	iu.mutation.SetIdentityToEnvironmentID(id)
	return iu
}

// SetNillableIdentityToEnvironmentID sets the "IdentityToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (iu *IdentityUpdate) SetNillableIdentityToEnvironmentID(id *uuid.UUID) *IdentityUpdate {
	if id != nil {
		iu = iu.SetIdentityToEnvironmentID(*id)
	}
	return iu
}

// SetIdentityToEnvironment sets the "IdentityToEnvironment" edge to the Environment entity.
func (iu *IdentityUpdate) SetIdentityToEnvironment(e *Environment) *IdentityUpdate {
	return iu.SetIdentityToEnvironmentID(e.ID)
}

// Mutation returns the IdentityMutation object of the builder.
func (iu *IdentityUpdate) Mutation() *IdentityMutation {
	return iu.mutation
}

// ClearIdentityToEnvironment clears the "IdentityToEnvironment" edge to the Environment entity.
func (iu *IdentityUpdate) ClearIdentityToEnvironment() *IdentityUpdate {
	iu.mutation.ClearIdentityToEnvironment()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IdentityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IdentityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IdentityUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IdentityUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IdentityUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *IdentityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   identity.Table,
			Columns: identity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: identity.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldHclID,
		})
	}
	if value, ok := iu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldFirstName,
		})
	}
	if value, ok := iu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldLastName,
		})
	}
	if value, ok := iu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldEmail,
		})
	}
	if value, ok := iu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldPassword,
		})
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldDescription,
		})
	}
	if value, ok := iu.mutation.AvatarFile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldAvatarFile,
		})
	}
	if value, ok := iu.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldVars,
		})
	}
	if value, ok := iu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldTags,
		})
	}
	if iu.mutation.IdentityToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: []string{identity.IdentityToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.IdentityToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: []string{identity.IdentityToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{identity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// IdentityUpdateOne is the builder for updating a single Identity entity.
type IdentityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IdentityMutation
}

// SetHclID sets the "hcl_id" field.
func (iuo *IdentityUpdateOne) SetHclID(s string) *IdentityUpdateOne {
	iuo.mutation.SetHclID(s)
	return iuo
}

// SetFirstName sets the "first_name" field.
func (iuo *IdentityUpdateOne) SetFirstName(s string) *IdentityUpdateOne {
	iuo.mutation.SetFirstName(s)
	return iuo
}

// SetLastName sets the "last_name" field.
func (iuo *IdentityUpdateOne) SetLastName(s string) *IdentityUpdateOne {
	iuo.mutation.SetLastName(s)
	return iuo
}

// SetEmail sets the "email" field.
func (iuo *IdentityUpdateOne) SetEmail(s string) *IdentityUpdateOne {
	iuo.mutation.SetEmail(s)
	return iuo
}

// SetPassword sets the "password" field.
func (iuo *IdentityUpdateOne) SetPassword(s string) *IdentityUpdateOne {
	iuo.mutation.SetPassword(s)
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *IdentityUpdateOne) SetDescription(s string) *IdentityUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetAvatarFile sets the "avatar_file" field.
func (iuo *IdentityUpdateOne) SetAvatarFile(s string) *IdentityUpdateOne {
	iuo.mutation.SetAvatarFile(s)
	return iuo
}

// SetVars sets the "vars" field.
func (iuo *IdentityUpdateOne) SetVars(m map[string]string) *IdentityUpdateOne {
	iuo.mutation.SetVars(m)
	return iuo
}

// SetTags sets the "tags" field.
func (iuo *IdentityUpdateOne) SetTags(m map[string]string) *IdentityUpdateOne {
	iuo.mutation.SetTags(m)
	return iuo
}

// SetIdentityToEnvironmentID sets the "IdentityToEnvironment" edge to the Environment entity by ID.
func (iuo *IdentityUpdateOne) SetIdentityToEnvironmentID(id uuid.UUID) *IdentityUpdateOne {
	iuo.mutation.SetIdentityToEnvironmentID(id)
	return iuo
}

// SetNillableIdentityToEnvironmentID sets the "IdentityToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (iuo *IdentityUpdateOne) SetNillableIdentityToEnvironmentID(id *uuid.UUID) *IdentityUpdateOne {
	if id != nil {
		iuo = iuo.SetIdentityToEnvironmentID(*id)
	}
	return iuo
}

// SetIdentityToEnvironment sets the "IdentityToEnvironment" edge to the Environment entity.
func (iuo *IdentityUpdateOne) SetIdentityToEnvironment(e *Environment) *IdentityUpdateOne {
	return iuo.SetIdentityToEnvironmentID(e.ID)
}

// Mutation returns the IdentityMutation object of the builder.
func (iuo *IdentityUpdateOne) Mutation() *IdentityMutation {
	return iuo.mutation
}

// ClearIdentityToEnvironment clears the "IdentityToEnvironment" edge to the Environment entity.
func (iuo *IdentityUpdateOne) ClearIdentityToEnvironment() *IdentityUpdateOne {
	iuo.mutation.ClearIdentityToEnvironment()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IdentityUpdateOne) Select(field string, fields ...string) *IdentityUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Identity entity.
func (iuo *IdentityUpdateOne) Save(ctx context.Context) (*Identity, error) {
	var (
		err  error
		node *Identity
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IdentityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IdentityUpdateOne) SaveX(ctx context.Context) *Identity {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IdentityUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IdentityUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *IdentityUpdateOne) sqlSave(ctx context.Context) (_node *Identity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   identity.Table,
			Columns: identity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: identity.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Identity.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, identity.FieldID)
		for _, f := range fields {
			if !identity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != identity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldHclID,
		})
	}
	if value, ok := iuo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldFirstName,
		})
	}
	if value, ok := iuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldLastName,
		})
	}
	if value, ok := iuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldEmail,
		})
	}
	if value, ok := iuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldPassword,
		})
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldDescription,
		})
	}
	if value, ok := iuo.mutation.AvatarFile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldAvatarFile,
		})
	}
	if value, ok := iuo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldVars,
		})
	}
	if value, ok := iuo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldTags,
		})
	}
	if iuo.mutation.IdentityToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: []string{identity.IdentityToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.IdentityToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: []string{identity.IdentityToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Identity{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{identity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}