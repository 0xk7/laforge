// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/plandiff"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/provisioningstep"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/google/uuid"
)

// PlanCreate is the builder for creating a Plan entity.
type PlanCreate struct {
	config
	mutation *PlanMutation
	hooks    []Hook
}

// SetStepNumber sets the "step_number" field.
func (pc *PlanCreate) SetStepNumber(i int) *PlanCreate {
	pc.mutation.SetStepNumber(i)
	return pc
}

// SetType sets the "type" field.
func (pc *PlanCreate) SetType(pl plan.Type) *PlanCreate {
	pc.mutation.SetType(pl)
	return pc
}

// SetBuildID sets the "build_id" field.
func (pc *PlanCreate) SetBuildID(s string) *PlanCreate {
	pc.mutation.SetBuildID(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PlanCreate) SetID(u uuid.UUID) *PlanCreate {
	pc.mutation.SetID(u)
	return pc
}

// AddPrevPlanIDs adds the "PrevPlan" edge to the Plan entity by IDs.
func (pc *PlanCreate) AddPrevPlanIDs(ids ...uuid.UUID) *PlanCreate {
	pc.mutation.AddPrevPlanIDs(ids...)
	return pc
}

// AddPrevPlan adds the "PrevPlan" edges to the Plan entity.
func (pc *PlanCreate) AddPrevPlan(p ...*Plan) *PlanCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPrevPlanIDs(ids...)
}

// AddNextPlanIDs adds the "NextPlan" edge to the Plan entity by IDs.
func (pc *PlanCreate) AddNextPlanIDs(ids ...uuid.UUID) *PlanCreate {
	pc.mutation.AddNextPlanIDs(ids...)
	return pc
}

// AddNextPlan adds the "NextPlan" edges to the Plan entity.
func (pc *PlanCreate) AddNextPlan(p ...*Plan) *PlanCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddNextPlanIDs(ids...)
}

// SetPlanToBuildID sets the "PlanToBuild" edge to the Build entity by ID.
func (pc *PlanCreate) SetPlanToBuildID(id uuid.UUID) *PlanCreate {
	pc.mutation.SetPlanToBuildID(id)
	return pc
}

// SetNillablePlanToBuildID sets the "PlanToBuild" edge to the Build entity by ID if the given value is not nil.
func (pc *PlanCreate) SetNillablePlanToBuildID(id *uuid.UUID) *PlanCreate {
	if id != nil {
		pc = pc.SetPlanToBuildID(*id)
	}
	return pc
}

// SetPlanToBuild sets the "PlanToBuild" edge to the Build entity.
func (pc *PlanCreate) SetPlanToBuild(b *Build) *PlanCreate {
	return pc.SetPlanToBuildID(b.ID)
}

// SetPlanToTeamID sets the "PlanToTeam" edge to the Team entity by ID.
func (pc *PlanCreate) SetPlanToTeamID(id uuid.UUID) *PlanCreate {
	pc.mutation.SetPlanToTeamID(id)
	return pc
}

// SetNillablePlanToTeamID sets the "PlanToTeam" edge to the Team entity by ID if the given value is not nil.
func (pc *PlanCreate) SetNillablePlanToTeamID(id *uuid.UUID) *PlanCreate {
	if id != nil {
		pc = pc.SetPlanToTeamID(*id)
	}
	return pc
}

// SetPlanToTeam sets the "PlanToTeam" edge to the Team entity.
func (pc *PlanCreate) SetPlanToTeam(t *Team) *PlanCreate {
	return pc.SetPlanToTeamID(t.ID)
}

// SetPlanToProvisionedNetworkID sets the "PlanToProvisionedNetwork" edge to the ProvisionedNetwork entity by ID.
func (pc *PlanCreate) SetPlanToProvisionedNetworkID(id uuid.UUID) *PlanCreate {
	pc.mutation.SetPlanToProvisionedNetworkID(id)
	return pc
}

// SetNillablePlanToProvisionedNetworkID sets the "PlanToProvisionedNetwork" edge to the ProvisionedNetwork entity by ID if the given value is not nil.
func (pc *PlanCreate) SetNillablePlanToProvisionedNetworkID(id *uuid.UUID) *PlanCreate {
	if id != nil {
		pc = pc.SetPlanToProvisionedNetworkID(*id)
	}
	return pc
}

// SetPlanToProvisionedNetwork sets the "PlanToProvisionedNetwork" edge to the ProvisionedNetwork entity.
func (pc *PlanCreate) SetPlanToProvisionedNetwork(p *ProvisionedNetwork) *PlanCreate {
	return pc.SetPlanToProvisionedNetworkID(p.ID)
}

// SetPlanToProvisionedHostID sets the "PlanToProvisionedHost" edge to the ProvisionedHost entity by ID.
func (pc *PlanCreate) SetPlanToProvisionedHostID(id uuid.UUID) *PlanCreate {
	pc.mutation.SetPlanToProvisionedHostID(id)
	return pc
}

// SetNillablePlanToProvisionedHostID sets the "PlanToProvisionedHost" edge to the ProvisionedHost entity by ID if the given value is not nil.
func (pc *PlanCreate) SetNillablePlanToProvisionedHostID(id *uuid.UUID) *PlanCreate {
	if id != nil {
		pc = pc.SetPlanToProvisionedHostID(*id)
	}
	return pc
}

// SetPlanToProvisionedHost sets the "PlanToProvisionedHost" edge to the ProvisionedHost entity.
func (pc *PlanCreate) SetPlanToProvisionedHost(p *ProvisionedHost) *PlanCreate {
	return pc.SetPlanToProvisionedHostID(p.ID)
}

// SetPlanToProvisioningStepID sets the "PlanToProvisioningStep" edge to the ProvisioningStep entity by ID.
func (pc *PlanCreate) SetPlanToProvisioningStepID(id uuid.UUID) *PlanCreate {
	pc.mutation.SetPlanToProvisioningStepID(id)
	return pc
}

// SetNillablePlanToProvisioningStepID sets the "PlanToProvisioningStep" edge to the ProvisioningStep entity by ID if the given value is not nil.
func (pc *PlanCreate) SetNillablePlanToProvisioningStepID(id *uuid.UUID) *PlanCreate {
	if id != nil {
		pc = pc.SetPlanToProvisioningStepID(*id)
	}
	return pc
}

// SetPlanToProvisioningStep sets the "PlanToProvisioningStep" edge to the ProvisioningStep entity.
func (pc *PlanCreate) SetPlanToProvisioningStep(p *ProvisioningStep) *PlanCreate {
	return pc.SetPlanToProvisioningStepID(p.ID)
}

// SetPlanToStatusID sets the "PlanToStatus" edge to the Status entity by ID.
func (pc *PlanCreate) SetPlanToStatusID(id uuid.UUID) *PlanCreate {
	pc.mutation.SetPlanToStatusID(id)
	return pc
}

// SetPlanToStatus sets the "PlanToStatus" edge to the Status entity.
func (pc *PlanCreate) SetPlanToStatus(s *Status) *PlanCreate {
	return pc.SetPlanToStatusID(s.ID)
}

// AddPlanToPlanDiffIDs adds the "PlanToPlanDiffs" edge to the PlanDiff entity by IDs.
func (pc *PlanCreate) AddPlanToPlanDiffIDs(ids ...uuid.UUID) *PlanCreate {
	pc.mutation.AddPlanToPlanDiffIDs(ids...)
	return pc
}

// AddPlanToPlanDiffs adds the "PlanToPlanDiffs" edges to the PlanDiff entity.
func (pc *PlanCreate) AddPlanToPlanDiffs(p ...*PlanDiff) *PlanCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPlanToPlanDiffIDs(ids...)
}

// Mutation returns the PlanMutation object of the builder.
func (pc *PlanCreate) Mutation() *PlanMutation {
	return pc.mutation
}

// Save creates the Plan in the database.
func (pc *PlanCreate) Save(ctx context.Context) (*Plan, error) {
	var (
		err  error
		node *Plan
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlanCreate) SaveX(ctx context.Context) *Plan {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlanCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlanCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlanCreate) defaults() {
	if _, ok := pc.mutation.ID(); !ok {
		v := plan.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlanCreate) check() error {
	if _, ok := pc.mutation.StepNumber(); !ok {
		return &ValidationError{Name: "step_number", err: errors.New(`ent: missing required field "step_number"`)}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := pc.mutation.GetType(); ok {
		if err := plan.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.BuildID(); !ok {
		return &ValidationError{Name: "build_id", err: errors.New(`ent: missing required field "build_id"`)}
	}
	if _, ok := pc.mutation.PlanToStatusID(); !ok {
		return &ValidationError{Name: "PlanToStatus", err: errors.New("ent: missing required edge \"PlanToStatus\"")}
	}
	return nil
}

func (pc *PlanCreate) sqlSave(ctx context.Context) (*Plan, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (pc *PlanCreate) createSpec() (*Plan, *sqlgraph.CreateSpec) {
	var (
		_node = &Plan{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: plan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: plan.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.StepNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: plan.FieldStepNumber,
		})
		_node.StepNumber = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: plan.FieldType,
		})
		_node.Type = value
	}
	if value, ok := pc.mutation.BuildID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plan.FieldBuildID,
		})
		_node.BuildID = value
	}
	if nodes := pc.mutation.PrevPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   plan.PrevPlanTable,
			Columns: plan.PrevPlanPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.NextPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.NextPlanTable,
			Columns: plan.NextPlanPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlanToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   plan.PlanToBuildTable,
			Columns: []string{plan.PlanToBuildColumn},
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
		_node.plan_plan_to_build = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlanToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.PlanToTeamTable,
			Columns: []string{plan.PlanToTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlanToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.PlanToProvisionedNetworkTable,
			Columns: []string{plan.PlanToProvisionedNetworkColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlanToProvisionedHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.PlanToProvisionedHostTable,
			Columns: []string{plan.PlanToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlanToProvisioningStepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.PlanToProvisioningStepTable,
			Columns: []string{plan.PlanToProvisioningStepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisioningstep.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlanToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.PlanToStatusTable,
			Columns: []string{plan.PlanToStatusColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlanToPlanDiffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   plan.PlanToPlanDiffsTable,
			Columns: []string{plan.PlanToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlanCreateBulk is the builder for creating many Plan entities in bulk.
type PlanCreateBulk struct {
	config
	builders []*PlanCreate
}

// Save creates the Plan entities in the database.
func (pcb *PlanCreateBulk) Save(ctx context.Context) ([]*Plan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Plan, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlanMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlanCreateBulk) SaveX(ctx context.Context) []*Plan {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlanCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlanCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}