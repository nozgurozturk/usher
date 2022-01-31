package query

import (
	"context"

	"github.com/nozgurozturk/usher/internal/domain/user"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

type GetUsersQuery struct{}

type GetUsersHandler struct {
	store *ent.Client
}

func NewGetUsersHandler(s *ent.Client) GetUsersHandler {
	return GetUsersHandler{
		store: s,
	}
}

func (h GetUsersHandler) Handle(q GetUsersQuery) ([]user.User, error) {
	ctx := context.Background()

	userDAOs, err := h.store.User.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	users := make([]user.User, len(userDAOs))
	for i := range userDAOs {
		users[i] = user.User{
			ID:   userDAOs[i].ID.String(),
			Name: userDAOs[i].Name,
		}
	}

	return users, nil
}
