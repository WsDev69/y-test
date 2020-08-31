package user

import (
	"context"
	img "image"
	"net/http"
	"y-test/internal/constant/model"
	"y-test/internal/module/image"
	"y-test/platform/minio"
	"y-test/platform/routers"
)

type Persistence interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	FindByID(ctx context.Context, userID string) (*model.User, error)
	Find(ctx context.Context, email, password string) (*model.User, error)
	FindPageable(ctx context.Context, pageSize, pageNumber int) ([]*model.User, error)
}

type Repository interface {
}

type Service interface {
	Create(ctx context.Context, u *model.User) (*model.User, error)
	Read(ctx context.Context, id string) (*model.User, error)
	ReadByCredentials(ctx context.Context, email, password string) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	UpdateAvatar(ctx context.Context, userid string, avatar img.Image) (*model.User, error)
	ReadPageable(ctx context.Context, pageNumber, pageSize int) ([]*model.User, error)
}

type service struct {
	p  Persistence
	os minio.ObjectStorage
	r  image.Service
}

// InitializeDomain is the function to initiate the business logic with services that'll be used by business logic
func NewUserService(p Persistence, os minio.ObjectStorage) Service {
	return &service{p: p, os: os, r: &image.Recize{}}
}

type Handler interface {
	SignUp(rw http.ResponseWriter, r *http.Request)
	Update(rw http.ResponseWriter, r *http.Request)
	Login(rw http.ResponseWriter, r *http.Request)
	Get(rw http.ResponseWriter, r *http.Request)
	GetPageable(rw http.ResponseWriter, r *http.Request)
	UpdateAvatar(rw http.ResponseWriter, r *http.Request)
}
type AuthHandler interface {
	Authentication(rw http.ResponseWriter, r *http.Request)
}

type Route interface {
	Routers() []*routers.Router
}
