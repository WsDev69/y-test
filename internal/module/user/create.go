package user

import (
	"context"
	"github.com/satori/go.uuid"
	"y-test/internal/constant/model"
)

func (s *service) Create(ctx context.Context, u *model.User) (*model.User, error) {
	id := uuid.NewV4().String()
	u.UserId = id
	up, e := s.p.Create(ctx, u)
	if e != nil {
		return nil, e
	}
	return up, nil
}
