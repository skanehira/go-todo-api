package main

import (
	"context"
	"todo/ent"
	"todo/ent/ogent"
)

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}

func (h handler) MarkDone(ctx context.Context, params ogent.MarkDoneParams) error {
	return h.client.Todo.UpdateOneID(params.ID).SetDone(true).Exec(ctx)
}
