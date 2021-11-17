// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/command"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/user"
	"github.com/google/uuid"
)

// CommandUpdate is the builder for updating Command entities.
type CommandUpdate struct {
	config
	hooks    []Hook
	mutation *CommandMutation
}

// Where appends a list predicates to the CommandUpdate builder.
func (cu *CommandUpdate) Where(ps ...predicate.Command) *CommandUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetHclID sets the "hcl_id" field.
func (cu *CommandUpdate) SetHclID(s string) *CommandUpdate {
	cu.mutation.SetHclID(s)
	return cu
}

// SetName sets the "name" field.
func (cu *CommandUpdate) SetName(s string) *CommandUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CommandUpdate) SetDescription(s string) *CommandUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetProgram sets the "program" field.
func (cu *CommandUpdate) SetProgram(s string) *CommandUpdate {
	cu.mutation.SetProgram(s)
	return cu
}

// SetArgs sets the "args" field.
func (cu *CommandUpdate) SetArgs(s []string) *CommandUpdate {
	cu.mutation.SetArgs(s)
	return cu
}

// SetIgnoreErrors sets the "ignore_errors" field.
func (cu *CommandUpdate) SetIgnoreErrors(b bool) *CommandUpdate {
	cu.mutation.SetIgnoreErrors(b)
	return cu
}

// SetDisabled sets the "disabled" field.
func (cu *CommandUpdate) SetDisabled(b bool) *CommandUpdate {
	cu.mutation.SetDisabled(b)
	return cu
}

// SetCooldown sets the "cooldown" field.
func (cu *CommandUpdate) SetCooldown(i int) *CommandUpdate {
	cu.mutation.ResetCooldown()
	cu.mutation.SetCooldown(i)
	return cu
}

// AddCooldown adds i to the "cooldown" field.
func (cu *CommandUpdate) AddCooldown(i int) *CommandUpdate {
	cu.mutation.AddCooldown(i)
	return cu
}

// SetTimeout sets the "timeout" field.
func (cu *CommandUpdate) SetTimeout(i int) *CommandUpdate {
	cu.mutation.ResetTimeout()
	cu.mutation.SetTimeout(i)
	return cu
}

// AddTimeout adds i to the "timeout" field.
func (cu *CommandUpdate) AddTimeout(i int) *CommandUpdate {
	cu.mutation.AddTimeout(i)
	return cu
}

// SetVars sets the "vars" field.
func (cu *CommandUpdate) SetVars(m map[string]string) *CommandUpdate {
	cu.mutation.SetVars(m)
	return cu
}

// SetTags sets the "tags" field.
func (cu *CommandUpdate) SetTags(m map[string]string) *CommandUpdate {
	cu.mutation.SetTags(m)
	return cu
}

// AddCommandToUserIDs adds the "CommandToUser" edge to the User entity by IDs.
func (cu *CommandUpdate) AddCommandToUserIDs(ids ...uuid.UUID) *CommandUpdate {
	cu.mutation.AddCommandToUserIDs(ids...)
	return cu
}

// AddCommandToUser adds the "CommandToUser" edges to the User entity.
func (cu *CommandUpdate) AddCommandToUser(u ...*User) *CommandUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddCommandToUserIDs(ids...)
}

// SetCommandToEnvironmentID sets the "CommandToEnvironment" edge to the Environment entity by ID.
func (cu *CommandUpdate) SetCommandToEnvironmentID(id uuid.UUID) *CommandUpdate {
	cu.mutation.SetCommandToEnvironmentID(id)
	return cu
}

// SetNillableCommandToEnvironmentID sets the "CommandToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (cu *CommandUpdate) SetNillableCommandToEnvironmentID(id *uuid.UUID) *CommandUpdate {
	if id != nil {
		cu = cu.SetCommandToEnvironmentID(*id)
	}
	return cu
}

// SetCommandToEnvironment sets the "CommandToEnvironment" edge to the Environment entity.
func (cu *CommandUpdate) SetCommandToEnvironment(e *Environment) *CommandUpdate {
	return cu.SetCommandToEnvironmentID(e.ID)
}

// Mutation returns the CommandMutation object of the builder.
func (cu *CommandUpdate) Mutation() *CommandMutation {
	return cu.mutation
}

// ClearCommandToUser clears all "CommandToUser" edges to the User entity.
func (cu *CommandUpdate) ClearCommandToUser() *CommandUpdate {
	cu.mutation.ClearCommandToUser()
	return cu
}

// RemoveCommandToUserIDs removes the "CommandToUser" edge to User entities by IDs.
func (cu *CommandUpdate) RemoveCommandToUserIDs(ids ...uuid.UUID) *CommandUpdate {
	cu.mutation.RemoveCommandToUserIDs(ids...)
	return cu
}

// RemoveCommandToUser removes "CommandToUser" edges to User entities.
func (cu *CommandUpdate) RemoveCommandToUser(u ...*User) *CommandUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveCommandToUserIDs(ids...)
}

// ClearCommandToEnvironment clears the "CommandToEnvironment" edge to the Environment entity.
func (cu *CommandUpdate) ClearCommandToEnvironment() *CommandUpdate {
	cu.mutation.ClearCommandToEnvironment()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommandUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommandUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommandUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommandUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommandUpdate) check() error {
	if v, ok := cu.mutation.Cooldown(); ok {
		if err := command.CooldownValidator(v); err != nil {
			return &ValidationError{Name: "cooldown", err: fmt.Errorf("ent: validator failed for field \"cooldown\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Timeout(); ok {
		if err := command.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	return nil
}

func (cu *CommandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   command.Table,
			Columns: command.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: command.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldHclID,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldName,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldDescription,
		})
	}
	if value, ok := cu.mutation.Program(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldProgram,
		})
	}
	if value, ok := cu.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldArgs,
		})
	}
	if value, ok := cu.mutation.IgnoreErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: command.FieldIgnoreErrors,
		})
	}
	if value, ok := cu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: command.FieldDisabled,
		})
	}
	if value, ok := cu.mutation.Cooldown(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldCooldown,
		})
	}
	if value, ok := cu.mutation.AddedCooldown(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldCooldown,
		})
	}
	if value, ok := cu.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldTimeout,
		})
	}
	if value, ok := cu.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldTimeout,
		})
	}
	if value, ok := cu.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldVars,
		})
	}
	if value, ok := cu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldTags,
		})
	}
	if cu.mutation.CommandToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToUserTable,
			Columns: []string{command.CommandToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCommandToUserIDs(); len(nodes) > 0 && !cu.mutation.CommandToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToUserTable,
			Columns: []string{command.CommandToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CommandToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToUserTable,
			Columns: []string{command.CommandToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CommandToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   command.CommandToEnvironmentTable,
			Columns: []string{command.CommandToEnvironmentColumn},
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
	if nodes := cu.mutation.CommandToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   command.CommandToEnvironmentTable,
			Columns: []string{command.CommandToEnvironmentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{command.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CommandUpdateOne is the builder for updating a single Command entity.
type CommandUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommandMutation
}

// SetHclID sets the "hcl_id" field.
func (cuo *CommandUpdateOne) SetHclID(s string) *CommandUpdateOne {
	cuo.mutation.SetHclID(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CommandUpdateOne) SetName(s string) *CommandUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CommandUpdateOne) SetDescription(s string) *CommandUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetProgram sets the "program" field.
func (cuo *CommandUpdateOne) SetProgram(s string) *CommandUpdateOne {
	cuo.mutation.SetProgram(s)
	return cuo
}

// SetArgs sets the "args" field.
func (cuo *CommandUpdateOne) SetArgs(s []string) *CommandUpdateOne {
	cuo.mutation.SetArgs(s)
	return cuo
}

// SetIgnoreErrors sets the "ignore_errors" field.
func (cuo *CommandUpdateOne) SetIgnoreErrors(b bool) *CommandUpdateOne {
	cuo.mutation.SetIgnoreErrors(b)
	return cuo
}

// SetDisabled sets the "disabled" field.
func (cuo *CommandUpdateOne) SetDisabled(b bool) *CommandUpdateOne {
	cuo.mutation.SetDisabled(b)
	return cuo
}

// SetCooldown sets the "cooldown" field.
func (cuo *CommandUpdateOne) SetCooldown(i int) *CommandUpdateOne {
	cuo.mutation.ResetCooldown()
	cuo.mutation.SetCooldown(i)
	return cuo
}

// AddCooldown adds i to the "cooldown" field.
func (cuo *CommandUpdateOne) AddCooldown(i int) *CommandUpdateOne {
	cuo.mutation.AddCooldown(i)
	return cuo
}

// SetTimeout sets the "timeout" field.
func (cuo *CommandUpdateOne) SetTimeout(i int) *CommandUpdateOne {
	cuo.mutation.ResetTimeout()
	cuo.mutation.SetTimeout(i)
	return cuo
}

// AddTimeout adds i to the "timeout" field.
func (cuo *CommandUpdateOne) AddTimeout(i int) *CommandUpdateOne {
	cuo.mutation.AddTimeout(i)
	return cuo
}

// SetVars sets the "vars" field.
func (cuo *CommandUpdateOne) SetVars(m map[string]string) *CommandUpdateOne {
	cuo.mutation.SetVars(m)
	return cuo
}

// SetTags sets the "tags" field.
func (cuo *CommandUpdateOne) SetTags(m map[string]string) *CommandUpdateOne {
	cuo.mutation.SetTags(m)
	return cuo
}

// AddCommandToUserIDs adds the "CommandToUser" edge to the User entity by IDs.
func (cuo *CommandUpdateOne) AddCommandToUserIDs(ids ...uuid.UUID) *CommandUpdateOne {
	cuo.mutation.AddCommandToUserIDs(ids...)
	return cuo
}

// AddCommandToUser adds the "CommandToUser" edges to the User entity.
func (cuo *CommandUpdateOne) AddCommandToUser(u ...*User) *CommandUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddCommandToUserIDs(ids...)
}

// SetCommandToEnvironmentID sets the "CommandToEnvironment" edge to the Environment entity by ID.
func (cuo *CommandUpdateOne) SetCommandToEnvironmentID(id uuid.UUID) *CommandUpdateOne {
	cuo.mutation.SetCommandToEnvironmentID(id)
	return cuo
}

// SetNillableCommandToEnvironmentID sets the "CommandToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (cuo *CommandUpdateOne) SetNillableCommandToEnvironmentID(id *uuid.UUID) *CommandUpdateOne {
	if id != nil {
		cuo = cuo.SetCommandToEnvironmentID(*id)
	}
	return cuo
}

// SetCommandToEnvironment sets the "CommandToEnvironment" edge to the Environment entity.
func (cuo *CommandUpdateOne) SetCommandToEnvironment(e *Environment) *CommandUpdateOne {
	return cuo.SetCommandToEnvironmentID(e.ID)
}

// Mutation returns the CommandMutation object of the builder.
func (cuo *CommandUpdateOne) Mutation() *CommandMutation {
	return cuo.mutation
}

// ClearCommandToUser clears all "CommandToUser" edges to the User entity.
func (cuo *CommandUpdateOne) ClearCommandToUser() *CommandUpdateOne {
	cuo.mutation.ClearCommandToUser()
	return cuo
}

// RemoveCommandToUserIDs removes the "CommandToUser" edge to User entities by IDs.
func (cuo *CommandUpdateOne) RemoveCommandToUserIDs(ids ...uuid.UUID) *CommandUpdateOne {
	cuo.mutation.RemoveCommandToUserIDs(ids...)
	return cuo
}

// RemoveCommandToUser removes "CommandToUser" edges to User entities.
func (cuo *CommandUpdateOne) RemoveCommandToUser(u ...*User) *CommandUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveCommandToUserIDs(ids...)
}

// ClearCommandToEnvironment clears the "CommandToEnvironment" edge to the Environment entity.
func (cuo *CommandUpdateOne) ClearCommandToEnvironment() *CommandUpdateOne {
	cuo.mutation.ClearCommandToEnvironment()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommandUpdateOne) Select(field string, fields ...string) *CommandUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Command entity.
func (cuo *CommandUpdateOne) Save(ctx context.Context) (*Command, error) {
	var (
		err  error
		node *Command
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommandUpdateOne) SaveX(ctx context.Context) *Command {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommandUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommandUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommandUpdateOne) check() error {
	if v, ok := cuo.mutation.Cooldown(); ok {
		if err := command.CooldownValidator(v); err != nil {
			return &ValidationError{Name: "cooldown", err: fmt.Errorf("ent: validator failed for field \"cooldown\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Timeout(); ok {
		if err := command.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	return nil
}

func (cuo *CommandUpdateOne) sqlSave(ctx context.Context) (_node *Command, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   command.Table,
			Columns: command.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: command.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Command.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, command.FieldID)
		for _, f := range fields {
			if !command.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != command.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldHclID,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldName,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.Program(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: command.FieldProgram,
		})
	}
	if value, ok := cuo.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldArgs,
		})
	}
	if value, ok := cuo.mutation.IgnoreErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: command.FieldIgnoreErrors,
		})
	}
	if value, ok := cuo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: command.FieldDisabled,
		})
	}
	if value, ok := cuo.mutation.Cooldown(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldCooldown,
		})
	}
	if value, ok := cuo.mutation.AddedCooldown(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldCooldown,
		})
	}
	if value, ok := cuo.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldTimeout,
		})
	}
	if value, ok := cuo.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: command.FieldTimeout,
		})
	}
	if value, ok := cuo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldVars,
		})
	}
	if value, ok := cuo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: command.FieldTags,
		})
	}
	if cuo.mutation.CommandToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToUserTable,
			Columns: []string{command.CommandToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCommandToUserIDs(); len(nodes) > 0 && !cuo.mutation.CommandToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToUserTable,
			Columns: []string{command.CommandToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CommandToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   command.CommandToUserTable,
			Columns: []string{command.CommandToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CommandToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   command.CommandToEnvironmentTable,
			Columns: []string{command.CommandToEnvironmentColumn},
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
	if nodes := cuo.mutation.CommandToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   command.CommandToEnvironmentTable,
			Columns: []string{command.CommandToEnvironmentColumn},
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
	_node = &Command{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{command.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}