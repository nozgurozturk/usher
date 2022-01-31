package query

import (
	"context"

	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

type GetLayoutsQuery struct{}

type GetLayoutsHandler struct {
	store *ent.Client
}

func NewGetLayoutsHandler(s *ent.Client) GetLayoutsHandler {
	return GetLayoutsHandler{
		store: s,
	}
}

func (h GetLayoutsHandler) Handle(q GetLayoutsQuery) ([]layout.Hall, error) {
	ctx := context.Background()

	layoutDAOs, err := h.store.Layout.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	layouts := make([]layout.Hall, len(layoutDAOs))
	for i := range layoutDAOs {
		layouts[i] = store.ToLayoutEntity(layoutDAOs[i])
	}

	return layouts, nil
}
