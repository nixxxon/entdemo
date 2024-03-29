// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/nixxxon/entdemo/ent/todohistory"
)

func (t *Todo) History() *TodoHistoryQuery {
	historyClient := NewTodoHistoryClient(t.config)
	return historyClient.Query().Where(todohistory.Ref(t.ID))
}

func (th *TodoHistory) Next(ctx context.Context) (*TodoHistory, error) {
	client := NewTodoHistoryClient(th.config)
	return client.Query().
		Where(
			todohistory.Ref(th.Ref),
			todohistory.HistoryTimeGT(th.HistoryTime),
		).
		Order(todohistory.ByHistoryTime()).
		First(ctx)
}

func (th *TodoHistory) Prev(ctx context.Context) (*TodoHistory, error) {
	client := NewTodoHistoryClient(th.config)
	return client.Query().
		Where(
			todohistory.Ref(th.Ref),
			todohistory.HistoryTimeLT(th.HistoryTime),
		).
		Order(todohistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (thq *TodoHistoryQuery) Earliest(ctx context.Context) (*TodoHistory, error) {
	return thq.
		Order(todohistory.ByHistoryTime()).
		First(ctx)
}

func (thq *TodoHistoryQuery) Latest(ctx context.Context) (*TodoHistory, error) {
	return thq.
		Order(todohistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (thq *TodoHistoryQuery) AsOf(ctx context.Context, time time.Time) (*TodoHistory, error) {
	return thq.
		Where(todohistory.HistoryTimeLTE(time)).
		Order(todohistory.ByHistoryTime(sql.OrderDesc())).
		First(ctx)
}

func (th *TodoHistory) Restore(ctx context.Context) (*Todo, error) {
	client := NewTodoClient(th.config)
	return client.
		UpdateOneID(th.Ref).
		SetOtherID(th.OtherID).
		SetName(th.Name).
		Save(ctx)
}
