package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	dao "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/layout"
)

type GetLayoutQuery struct {
	LayoutID string
}

type GetLayoutHandler struct {
	store *ent.Client
}

func NewGetLayoutHandler(s *ent.Client) GetLayoutHandler {
	return GetLayoutHandler{
		store: s,
	}
}

func (h GetLayoutHandler) Handle(q GetLayoutQuery) (layout.Hall, error) {
	ctx := context.Background()

	layoutID, err := uuid.Parse(q.LayoutID)
	if err != nil {
		return nil, err
	}

	layoutDAOs, err := h.store.Layout.Query().Where(dao.ID(layoutID)).WithSections(
		func(q *ent.SectionQuery) {
			q.WithRows(func(rq *ent.RowQuery) {
				rq.WithSeats()
			})
		},
	).Only(ctx)
	if err != nil {
		return nil, err
	}

	l := store.ToLayoutEntity(layoutDAOs)

	return l, nil
}
