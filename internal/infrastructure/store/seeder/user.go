package seeder

import (
	"context"

	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

func CreateUsers(c *ent.Client, users ...string) []*ent.User {
	ctx := context.Background()

	userCreates := make([]*ent.UserCreate, len(users))

	for i, u := range users {
		userCreates[i] = c.User.Create().SetName(u)
	}

	return c.User.CreateBulk(userCreates...).SaveX(ctx)
}
