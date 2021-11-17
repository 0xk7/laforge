// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/google/uuid"
)

// TeamUpdate is the builder for updating Team entities.
type TeamUpdate struct {
	config
	hooks    []Hook
	mutation *TeamMutation
}

// Where appends a list predicates to the TeamUpdate builder.
func (tu *TeamUpdate) Where(ps ...predicate.Team) *TeamUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTeamNumber sets the "team_number" field.
func (tu *TeamUpdate) SetTeamNumber(i int) *TeamUpdate {
	tu.mutation.ResetTeamNumber()
	tu.mutation.SetTeamNumber(i)
	return tu
}

// AddTeamNumber adds i to the "team_number" field.
func (tu *TeamUpdate) AddTeamNumber(i int) *TeamUpdate {
	tu.mutation.AddTeamNumber(i)
	return tu
}

// SetVars sets the "vars" field.
func (tu *TeamUpdate) SetVars(m map[string]string) *TeamUpdate {
	tu.mutation.SetVars(m)
	return tu
}

// SetTeamToBuildID sets the "TeamToBuild" edge to the Build entity by ID.
func (tu *TeamUpdate) SetTeamToBuildID(id uuid.UUID) *TeamUpdate {
	tu.mutation.SetTeamToBuildID(id)
	return tu
}

// SetTeamToBuild sets the "TeamToBuild" edge to the Build entity.
func (tu *TeamUpdate) SetTeamToBuild(b *Build) *TeamUpdate {
	return tu.SetTeamToBuildID(b.ID)
}

// SetTeamToStatusID sets the "TeamToStatus" edge to the Status entity by ID.
func (tu *TeamUpdate) SetTeamToStatusID(id uuid.UUID) *TeamUpdate {
	tu.mutation.SetTeamToStatusID(id)
	return tu
}

// SetNillableTeamToStatusID sets the "TeamToStatus" edge to the Status entity by ID if the given value is not nil.
func (tu *TeamUpdate) SetNillableTeamToStatusID(id *uuid.UUID) *TeamUpdate {
	if id != nil {
		tu = tu.SetTeamToStatusID(*id)
	}
	return tu
}

// SetTeamToStatus sets the "TeamToStatus" edge to the Status entity.
func (tu *TeamUpdate) SetTeamToStatus(s *Status) *TeamUpdate {
	return tu.SetTeamToStatusID(s.ID)
}

// AddTeamToProvisionedNetworkIDs adds the "TeamToProvisionedNetwork" edge to the ProvisionedNetwork entity by IDs.
func (tu *TeamUpdate) AddTeamToProvisionedNetworkIDs(ids ...uuid.UUID) *TeamUpdate {
	tu.mutation.AddTeamToProvisionedNetworkIDs(ids...)
	return tu
}

// AddTeamToProvisionedNetwork adds the "TeamToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (tu *TeamUpdate) AddTeamToProvisionedNetwork(p ...*ProvisionedNetwork) *TeamUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.AddTeamToProvisionedNetworkIDs(ids...)
}

// SetTeamToPlanID sets the "TeamToPlan" edge to the Plan entity by ID.
func (tu *TeamUpdate) SetTeamToPlanID(id uuid.UUID) *TeamUpdate {
	tu.mutation.SetTeamToPlanID(id)
	return tu
}

// SetNillableTeamToPlanID sets the "TeamToPlan" edge to the Plan entity by ID if the given value is not nil.
func (tu *TeamUpdate) SetNillableTeamToPlanID(id *uuid.UUID) *TeamUpdate {
	if id != nil {
		tu = tu.SetTeamToPlanID(*id)
	}
	return tu
}

// SetTeamToPlan sets the "TeamToPlan" edge to the Plan entity.
func (tu *TeamUpdate) SetTeamToPlan(p *Plan) *TeamUpdate {
	return tu.SetTeamToPlanID(p.ID)
}

// Mutation returns the TeamMutation object of the builder.
func (tu *TeamUpdate) Mutation() *TeamMutation {
	return tu.mutation
}

// ClearTeamToBuild clears the "TeamToBuild" edge to the Build entity.
func (tu *TeamUpdate) ClearTeamToBuild() *TeamUpdate {
	tu.mutation.ClearTeamToBuild()
	return tu
}

// ClearTeamToStatus clears the "TeamToStatus" edge to the Status entity.
func (tu *TeamUpdate) ClearTeamToStatus() *TeamUpdate {
	tu.mutation.ClearTeamToStatus()
	return tu
}

// ClearTeamToProvisionedNetwork clears all "TeamToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (tu *TeamUpdate) ClearTeamToProvisionedNetwork() *TeamUpdate {
	tu.mutation.ClearTeamToProvisionedNetwork()
	return tu
}

// RemoveTeamToProvisionedNetworkIDs removes the "TeamToProvisionedNetwork" edge to ProvisionedNetwork entities by IDs.
func (tu *TeamUpdate) RemoveTeamToProvisionedNetworkIDs(ids ...uuid.UUID) *TeamUpdate {
	tu.mutation.RemoveTeamToProvisionedNetworkIDs(ids...)
	return tu
}

// RemoveTeamToProvisionedNetwork removes "TeamToProvisionedNetwork" edges to ProvisionedNetwork entities.
func (tu *TeamUpdate) RemoveTeamToProvisionedNetwork(p ...*ProvisionedNetwork) *TeamUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.RemoveTeamToProvisionedNetworkIDs(ids...)
}

// ClearTeamToPlan clears the "TeamToPlan" edge to the Plan entity.
func (tu *TeamUpdate) ClearTeamToPlan() *TeamUpdate {
	tu.mutation.ClearTeamToPlan()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TeamUpdate) check() error {
	if _, ok := tu.mutation.TeamToBuildID(); tu.mutation.TeamToBuildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"TeamToBuild\"")
	}
	return nil
}

func (tu *TeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   team.Table,
			Columns: team.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: team.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.TeamNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: team.FieldTeamNumber,
		})
	}
	if value, ok := tu.mutation.AddedTeamNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: team.FieldTeamNumber,
		})
	}
	if value, ok := tu.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: team.FieldVars,
		})
	}
	if tu.mutation.TeamToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   team.TeamToBuildTable,
			Columns: []string{team.TeamToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   team.TeamToBuildTable,
			Columns: []string{team.TeamToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TeamToStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   team.TeamToStatusTable,
			Columns: []string{team.TeamToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   team.TeamToStatusTable,
			Columns: []string{team.TeamToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TeamToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.TeamToProvisionedNetworkTable,
			Columns: []string{team.TeamToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTeamToProvisionedNetworkIDs(); len(nodes) > 0 && !tu.mutation.TeamToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.TeamToProvisionedNetworkTable,
			Columns: []string{team.TeamToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.TeamToProvisionedNetworkTable,
			Columns: []string{team.TeamToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TeamToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   team.TeamToPlanTable,
			Columns: []string{team.TeamToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   team.TeamToPlanTable,
			Columns: []string{team.TeamToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TeamUpdateOne is the builder for updating a single Team entity.
type TeamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamMutation
}

// SetTeamNumber sets the "team_number" field.
func (tuo *TeamUpdateOne) SetTeamNumber(i int) *TeamUpdateOne {
	tuo.mutation.ResetTeamNumber()
	tuo.mutation.SetTeamNumber(i)
	return tuo
}

// AddTeamNumber adds i to the "team_number" field.
func (tuo *TeamUpdateOne) AddTeamNumber(i int) *TeamUpdateOne {
	tuo.mutation.AddTeamNumber(i)
	return tuo
}

// SetVars sets the "vars" field.
func (tuo *TeamUpdateOne) SetVars(m map[string]string) *TeamUpdateOne {
	tuo.mutation.SetVars(m)
	return tuo
}

// SetTeamToBuildID sets the "TeamToBuild" edge to the Build entity by ID.
func (tuo *TeamUpdateOne) SetTeamToBuildID(id uuid.UUID) *TeamUpdateOne {
	tuo.mutation.SetTeamToBuildID(id)
	return tuo
}

// SetTeamToBuild sets the "TeamToBuild" edge to the Build entity.
func (tuo *TeamUpdateOne) SetTeamToBuild(b *Build) *TeamUpdateOne {
	return tuo.SetTeamToBuildID(b.ID)
}

// SetTeamToStatusID sets the "TeamToStatus" edge to the Status entity by ID.
func (tuo *TeamUpdateOne) SetTeamToStatusID(id uuid.UUID) *TeamUpdateOne {
	tuo.mutation.SetTeamToStatusID(id)
	return tuo
}

// SetNillableTeamToStatusID sets the "TeamToStatus" edge to the Status entity by ID if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableTeamToStatusID(id *uuid.UUID) *TeamUpdateOne {
	if id != nil {
		tuo = tuo.SetTeamToStatusID(*id)
	}
	return tuo
}

// SetTeamToStatus sets the "TeamToStatus" edge to the Status entity.
func (tuo *TeamUpdateOne) SetTeamToStatus(s *Status) *TeamUpdateOne {
	return tuo.SetTeamToStatusID(s.ID)
}

// AddTeamToProvisionedNetworkIDs adds the "TeamToProvisionedNetwork" edge to the ProvisionedNetwork entity by IDs.
func (tuo *TeamUpdateOne) AddTeamToProvisionedNetworkIDs(ids ...uuid.UUID) *TeamUpdateOne {
	tuo.mutation.AddTeamToProvisionedNetworkIDs(ids...)
	return tuo
}

// AddTeamToProvisionedNetwork adds the "TeamToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (tuo *TeamUpdateOne) AddTeamToProvisionedNetwork(p ...*ProvisionedNetwork) *TeamUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.AddTeamToProvisionedNetworkIDs(ids...)
}

// SetTeamToPlanID sets the "TeamToPlan" edge to the Plan entity by ID.
func (tuo *TeamUpdateOne) SetTeamToPlanID(id uuid.UUID) *TeamUpdateOne {
	tuo.mutation.SetTeamToPlanID(id)
	return tuo
}

// SetNillableTeamToPlanID sets the "TeamToPlan" edge to the Plan entity by ID if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableTeamToPlanID(id *uuid.UUID) *TeamUpdateOne {
	if id != nil {
		tuo = tuo.SetTeamToPlanID(*id)
	}
	return tuo
}

// SetTeamToPlan sets the "TeamToPlan" edge to the Plan entity.
func (tuo *TeamUpdateOne) SetTeamToPlan(p *Plan) *TeamUpdateOne {
	return tuo.SetTeamToPlanID(p.ID)
}

// Mutation returns the TeamMutation object of the builder.
func (tuo *TeamUpdateOne) Mutation() *TeamMutation {
	return tuo.mutation
}

// ClearTeamToBuild clears the "TeamToBuild" edge to the Build entity.
func (tuo *TeamUpdateOne) ClearTeamToBuild() *TeamUpdateOne {
	tuo.mutation.ClearTeamToBuild()
	return tuo
}

// ClearTeamToStatus clears the "TeamToStatus" edge to the Status entity.
func (tuo *TeamUpdateOne) ClearTeamToStatus() *TeamUpdateOne {
	tuo.mutation.ClearTeamToStatus()
	return tuo
}

// ClearTeamToProvisionedNetwork clears all "TeamToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (tuo *TeamUpdateOne) ClearTeamToProvisionedNetwork() *TeamUpdateOne {
	tuo.mutation.ClearTeamToProvisionedNetwork()
	return tuo
}

// RemoveTeamToProvisionedNetworkIDs removes the "TeamToProvisionedNetwork" edge to ProvisionedNetwork entities by IDs.
func (tuo *TeamUpdateOne) RemoveTeamToProvisionedNetworkIDs(ids ...uuid.UUID) *TeamUpdateOne {
	tuo.mutation.RemoveTeamToProvisionedNetworkIDs(ids...)
	return tuo
}

// RemoveTeamToProvisionedNetwork removes "TeamToProvisionedNetwork" edges to ProvisionedNetwork entities.
func (tuo *TeamUpdateOne) RemoveTeamToProvisionedNetwork(p ...*ProvisionedNetwork) *TeamUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.RemoveTeamToProvisionedNetworkIDs(ids...)
}

// ClearTeamToPlan clears the "TeamToPlan" edge to the Plan entity.
func (tuo *TeamUpdateOne) ClearTeamToPlan() *TeamUpdateOne {
	tuo.mutation.ClearTeamToPlan()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamUpdateOne) Select(field string, fields ...string) *TeamUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Team entity.
func (tuo *TeamUpdateOne) Save(ctx context.Context) (*Team, error) {
	var (
		err  error
		node *Team
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamUpdateOne) SaveX(ctx context.Context) *Team {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TeamUpdateOne) check() error {
	if _, ok := tuo.mutation.TeamToBuildID(); tuo.mutation.TeamToBuildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"TeamToBuild\"")
	}
	return nil
}

func (tuo *TeamUpdateOne) sqlSave(ctx context.Context) (_node *Team, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   team.Table,
			Columns: team.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: team.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Team.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for _, f := range fields {
			if !team.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.TeamNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: team.FieldTeamNumber,
		})
	}
	if value, ok := tuo.mutation.AddedTeamNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: team.FieldTeamNumber,
		})
	}
	if value, ok := tuo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: team.FieldVars,
		})
	}
	if tuo.mutation.TeamToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   team.TeamToBuildTable,
			Columns: []string{team.TeamToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   team.TeamToBuildTable,
			Columns: []string{team.TeamToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TeamToStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   team.TeamToStatusTable,
			Columns: []string{team.TeamToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   team.TeamToStatusTable,
			Columns: []string{team.TeamToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TeamToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.TeamToProvisionedNetworkTable,
			Columns: []string{team.TeamToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTeamToProvisionedNetworkIDs(); len(nodes) > 0 && !tuo.mutation.TeamToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.TeamToProvisionedNetworkTable,
			Columns: []string{team.TeamToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.TeamToProvisionedNetworkTable,
			Columns: []string{team.TeamToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TeamToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   team.TeamToPlanTable,
			Columns: []string{team.TeamToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   team.TeamToPlanTable,
			Columns: []string{team.TeamToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Team{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}