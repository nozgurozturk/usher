package command

import (
	"context"

	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

type ResetDatabaseCommand struct{}

type ResetDatabaseHandler struct {
	store *ent.Client
}

func NewResetDatabaseHandler(s *ent.Client) ResetDatabaseHandler {
	return ResetDatabaseHandler{
		store: s,
	}
}

func (h ResetDatabaseHandler) Handle(c ResetDatabaseCommand) {
	ctx := context.Background()
	store.ResetDatabase(ctx, h.store)
	store.CreateInitialData(ctx, h.store)
}
