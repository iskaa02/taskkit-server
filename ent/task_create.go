// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iskaa02/taskkit-server/ent/list"
	"github.com/iskaa02/taskkit-server/ent/task"
	"github.com/iskaa02/taskkit-server/types"
	"gopkg.in/guregu/null.v4"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TaskCreate) SetName(s string) *TaskCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetListID sets the "list_id" field.
func (tc *TaskCreate) SetListID(s string) *TaskCreate {
	tc.mutation.SetListID(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TaskCreate) SetDescription(n null.String) *TaskCreate {
	tc.mutation.SetDescription(n)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDescription(n *null.String) *TaskCreate {
	if n != nil {
		tc.SetDescription(*n)
	}
	return tc
}

// SetReminder sets the "reminder" field.
func (tc *TaskCreate) SetReminder(n null.Time) *TaskCreate {
	tc.mutation.SetReminder(n)
	return tc
}

// SetNillableReminder sets the "reminder" field if the given value is not nil.
func (tc *TaskCreate) SetNillableReminder(n *null.Time) *TaskCreate {
	if n != nil {
		tc.SetReminder(*n)
	}
	return tc
}

// SetRepeat sets the "repeat" field.
func (tc *TaskCreate) SetRepeat(n null.String) *TaskCreate {
	tc.mutation.SetRepeat(n)
	return tc
}

// SetNillableRepeat sets the "repeat" field if the given value is not nil.
func (tc *TaskCreate) SetNillableRepeat(n *null.String) *TaskCreate {
	if n != nil {
		tc.SetRepeat(*n)
	}
	return tc
}

// SetSubtasks sets the "subtasks" field.
func (tc *TaskCreate) SetSubtasks(t *types.Subtasks) *TaskCreate {
	tc.mutation.SetSubtasks(t)
	return tc
}

// SetIsCompleted sets the "is_completed" field.
func (tc *TaskCreate) SetIsCompleted(b bool) *TaskCreate {
	tc.mutation.SetIsCompleted(b)
	return tc
}

// SetNillableIsCompleted sets the "is_completed" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIsCompleted(b *bool) *TaskCreate {
	if b != nil {
		tc.SetIsCompleted(*b)
	}
	return tc
}

// SetIsDeleted sets the "is_deleted" field.
func (tc *TaskCreate) SetIsDeleted(b bool) *TaskCreate {
	tc.mutation.SetIsDeleted(b)
	return tc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIsDeleted(b *bool) *TaskCreate {
	if b != nil {
		tc.SetIsDeleted(*b)
	}
	return tc
}

// SetLastModified sets the "last_modified" field.
func (tc *TaskCreate) SetLastModified(t time.Time) *TaskCreate {
	tc.mutation.SetLastModified(t)
	return tc
}

// SetNillableLastModified sets the "last_modified" field if the given value is not nil.
func (tc *TaskCreate) SetNillableLastModified(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetLastModified(*t)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TaskCreate) SetID(s string) *TaskCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetList sets the "list" edge to the List entity.
func (tc *TaskCreate) SetList(l *List) *TaskCreate {
	return tc.SetListID(l.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.IsCompleted(); !ok {
		v := task.DefaultIsCompleted
		tc.mutation.SetIsCompleted(v)
	}
	if _, ok := tc.mutation.IsDeleted(); !ok {
		v := task.DefaultIsDeleted
		tc.mutation.SetIsDeleted(v)
	}
	if _, ok := tc.mutation.LastModified(); !ok {
		v := task.DefaultLastModified()
		tc.mutation.SetLastModified(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Task.name"`)}
	}
	if _, ok := tc.mutation.ListID(); !ok {
		return &ValidationError{Name: "list_id", err: errors.New(`ent: missing required field "Task.list_id"`)}
	}
	if _, ok := tc.mutation.Subtasks(); !ok {
		return &ValidationError{Name: "subtasks", err: errors.New(`ent: missing required field "Task.subtasks"`)}
	}
	if _, ok := tc.mutation.IsCompleted(); !ok {
		return &ValidationError{Name: "is_completed", err: errors.New(`ent: missing required field "Task.is_completed"`)}
	}
	if _, ok := tc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "Task.is_deleted"`)}
	}
	if _, ok := tc.mutation.LastModified(); !ok {
		return &ValidationError{Name: "last_modified", err: errors.New(`ent: missing required field "Task.last_modified"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := tc.mutation.ListID(); !ok {
		return &ValidationError{Name: "list", err: errors.New(`ent: missing required edge "Task.list"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Task.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: task.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := tc.mutation.Reminder(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldReminder,
		})
		_node.Reminder = value
	}
	if value, ok := tc.mutation.Repeat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldRepeat,
		})
		_node.Repeat = value
	}
	if value, ok := tc.mutation.Subtasks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldSubtasks,
		})
		_node.Subtasks = value
	}
	if value, ok := tc.mutation.IsCompleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldIsCompleted,
		})
		_node.IsCompleted = value
	}
	if value, ok := tc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := tc.mutation.LastModified(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldLastModified,
		})
		_node.LastModified = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := tc.mutation.ListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.ListTable,
			Columns: []string{task.ListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: list.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ListID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
