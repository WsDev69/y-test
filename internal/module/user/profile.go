package user

import (
	"context"
	"y-test/internal/constant/model"
)

func (s *service) Read(ctx context.Context, id string) (*model.User, error) {
	return s.p.FindByID(ctx, id)
}

func (s *service) ReadByCredentials(ctx context.Context, email, password string) (*model.User, error) {
	return s.p.Find(ctx, email, password)
}

func (s *service) ReadPageable(ctx context.Context, pageNumber, pageSize int) ([]*model.User, error)  {
	return s.p.FindPageable(ctx, pageNumber, pageSize)
}
