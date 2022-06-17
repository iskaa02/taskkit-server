// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iskaa02/taskkit-server/ent/list"
	"github.com/iskaa02/taskkit-server/ent/predicate"
	"github.com/iskaa02/taskkit-server/ent/task"
	"github.com/iskaa02/taskkit-server/ent/theme"
)

// ListUpdate is the builder for updating List entities.
type ListUpdate struct {
	config
	hooks    []Hook
	mutation *ListMutation
}

// Where appends a list predicates to the ListUpdate builder.
func (lu *ListUpdate) Where(ps ...predicate.List) *ListUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *ListUpdate) SetName(s string) *ListUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetThemeID sets the "theme_id" field.
func (lu *ListUpdate) SetThemeID(i int64) *ListUpdate {
	lu.mutation.SetThemeID(i)
	return lu
}

// SetIsDeleted sets the "is_deleted" field.
func (lu *ListUpdate) SetIsDeleted(b bool) *ListUpdate {
	lu.mutation.SetIsDeleted(b)
	return lu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (lu *ListUpdate) SetNillableIsDeleted(b *bool) *ListUpdate {
	if b != nil {
		lu.SetIsDeleted(*b)
	}
	return lu
}

// SetLastModified sets the "last_modified" field.
func (lu *ListUpdate) SetLastModified(t time.Time) *ListUpdate {
	lu.mutation.SetLastModified(t)
	return lu
}

// SetNillableLastModified sets the "last_modified" field if the given value is not nil.
func (lu *ListUpdate) SetNillableLastModified(t *time.Time) *ListUpdate {
	if t != nil {
		lu.SetLastModified(*t)
	}
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *ListUpdate) SetCreatedAt(t time.Time) *ListUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *ListUpdate) SetNillableCreatedAt(t *time.Time) *ListUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetTheme sets the "theme" edge to the Theme entity.
func (lu *ListUpdate) SetTheme(t *Theme) *ListUpdate {
	return lu.SetThemeID(t.ID)
}

// AddTaskIDs adds the "task" edge to the Task entity by IDs.
func (lu *ListUpdate) AddTaskIDs(ids ...string) *ListUpdate {
	lu.mutation.AddTaskIDs(ids...)
	return lu
}

// AddTask adds the "task" edges to the Task entity.
func (lu *ListUpdate) AddTask(t ...*Task) *ListUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return lu.AddTaskIDs(ids...)
}

// Mutation returns the ListMutation object of the builder.
func (lu *ListUpdate) Mutation() *ListMutation {
	return lu.mutation
}

// ClearTheme clears the "theme" edge to the Theme entity.
func (lu *ListUpdate) ClearTheme() *ListUpdate {
	lu.mutation.ClearTheme()
	return lu
}

// ClearTask clears all "task" edges to the Task entity.
func (lu *ListUpdate) ClearTask() *ListUpdate {
	lu.mutation.ClearTask()
	return lu
}

// RemoveTaskIDs removes the "task" edge to Task entities by IDs.
func (lu *ListUpdate) RemoveTaskIDs(ids ...string) *ListUpdate {
	lu.mutation.RemoveTaskIDs(ids...)
	return lu
}

// RemoveTask removes "task" edges to Task entities.
func (lu *ListUpdate) RemoveTask(t ...*Task) *ListUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return lu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *ListUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *ListUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *ListUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *ListUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *ListUpdate) check() error {
	if _, ok := lu.mutation.ThemeID(); lu.mutation.ThemeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "List.theme"`)
	}
	return nil
}

func (lu *ListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   list.Table,
			Columns: list.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: list.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: list.FieldName,
		})
	}
	if value, ok := lu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: list.FieldIsDeleted,
		})
	}
	if value, ok := lu.mutation.LastModified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: list.FieldLastModified,
		})
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: list.FieldCreatedAt,
		})
	}
	if lu.mutation.ThemeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   list.ThemeTable,
			Columns: []string{list.ThemeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: theme.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ThemeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   list.ThemeTable,
			Columns: []string{list.ThemeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: theme.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   list.TaskTable,
			Columns: []string{list.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedTaskIDs(); len(nodes) > 0 && !lu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   list.TaskTable,
			Columns: []string{list.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   list.TaskTable,
			Columns: []string{list.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{list.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ListUpdateOne is the builder for updating a single List entity.
type ListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ListMutation
}

// SetName sets the "name" field.
func (luo *ListUpdateOne) SetName(s string) *ListUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetThemeID sets the "theme_id" field.
func (luo *ListUpdateOne) SetThemeID(i int64) *ListUpdateOne {
	luo.mutation.SetThemeID(i)
	return luo
}

// SetIsDeleted sets the "is_deleted" field.
func (luo *ListUpdateOne) SetIsDeleted(b bool) *ListUpdateOne {
	luo.mutation.SetIsDeleted(b)
	return luo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (luo *ListUpdateOne) SetNillableIsDeleted(b *bool) *ListUpdateOne {
	if b != nil {
		luo.SetIsDeleted(*b)
	}
	return luo
}

// SetLastModified sets the "last_modified" field.
func (luo *ListUpdateOne) SetLastModified(t time.Time) *ListUpdateOne {
	luo.mutation.SetLastModified(t)
	return luo
}

// SetNillableLastModified sets the "last_modified" field if the given value is not nil.
func (luo *ListUpdateOne) SetNillableLastModified(t *time.Time) *ListUpdateOne {
	if t != nil {
		luo.SetLastModified(*t)
	}
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *ListUpdateOne) SetCreatedAt(t time.Time) *ListUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *ListUpdateOne) SetNillableCreatedAt(t *time.Time) *ListUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetTheme sets the "theme" edge to the Theme entity.
func (luo *ListUpdateOne) SetTheme(t *Theme) *ListUpdateOne {
	return luo.SetThemeID(t.ID)
}

// AddTaskIDs adds the "task" edge to the Task entity by IDs.
func (luo *ListUpdateOne) AddTaskIDs(ids ...string) *ListUpdateOne {
	luo.mutation.AddTaskIDs(ids...)
	return luo
}

// AddTask adds the "task" edges to the Task entity.
func (luo *ListUpdateOne) AddTask(t ...*Task) *ListUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return luo.AddTaskIDs(ids...)
}

// Mutation returns the ListMutation object of the builder.
func (luo *ListUpdateOne) Mutation() *ListMutation {
	return luo.mutation
}

// ClearTheme clears the "theme" edge to the Theme entity.
func (luo *ListUpdateOne) ClearTheme() *ListUpdateOne {
	luo.mutation.ClearTheme()
	return luo
}

// ClearTask clears all "task" edges to the Task entity.
func (luo *ListUpdateOne) ClearTask() *ListUpdateOne {
	luo.mutation.ClearTask()
	return luo
}

// RemoveTaskIDs removes the "task" edge to Task entities by IDs.
func (luo *ListUpdateOne) RemoveTaskIDs(ids ...string) *ListUpdateOne {
	luo.mutation.RemoveTaskIDs(ids...)
	return luo
}

// RemoveTask removes "task" edges to Task entities.
func (luo *ListUpdateOne) RemoveTask(t ...*Task) *ListUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return luo.RemoveTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *ListUpdateOne) Select(field string, fields ...string) *ListUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated List entity.
func (luo *ListUpdateOne) Save(ctx context.Context) (*List, error) {
	var (
		err  error
		node *List
	)
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *ListUpdateOne) SaveX(ctx context.Context) *List {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *ListUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *ListUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *ListUpdateOne) check() error {
	if _, ok := luo.mutation.ThemeID(); luo.mutation.ThemeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "List.theme"`)
	}
	return nil
}

func (luo *ListUpdateOne) sqlSave(ctx context.Context) (_node *List, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   list.Table,
			Columns: list.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: list.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "List.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, list.FieldID)
		for _, f := range fields {
			if !list.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != list.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: list.FieldName,
		})
	}
	if value, ok := luo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: list.FieldIsDeleted,
		})
	}
	if value, ok := luo.mutation.LastModified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: list.FieldLastModified,
		})
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: list.FieldCreatedAt,
		})
	}
	if luo.mutation.ThemeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   list.ThemeTable,
			Columns: []string{list.ThemeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: theme.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ThemeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   list.ThemeTable,
			Columns: []string{list.ThemeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: theme.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   list.TaskTable,
			Columns: []string{list.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedTaskIDs(); len(nodes) > 0 && !luo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   list.TaskTable,
			Columns: []string{list.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   list.TaskTable,
			Columns: []string{list.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &List{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{list.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
