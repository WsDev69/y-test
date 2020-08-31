package routing

import (
	"net/http"
	"y-test/internal/module/user"
	"y-test/platform/routers"
)

type userHandlers struct {
	handler user.Handler
}

func NewUserHandlers(handler user.Handler) *userHandlers {
	return &userHandlers{handler: handler}
}

func (uh *userHandlers) Routers() []*routers.Router {
	return []*routers.Router{
		{
			Method:  http.MethodPost,
			URL:     "/api/v1/user/signup",
			Handler: toHandler(uh.handler.SignUp),
			Open:    true,
		},
		{
			Method:  http.MethodPost,
			URL:     "/api/v1/user/login",
			Handler: toHandler(uh.handler.Login),
			Open:    true,
		},
		{
			Method:  http.MethodPatch,
			URL:     "/api/v1/user/{id}/update",
			Handler: toHandler(uh.handler.Update),
			Open:    false,
		},
		{
			Method:  http.MethodGet,
			URL:     "/api/v1/user/{id}",
			Handler: toHandler(uh.handler.Get),
			Open:    false,
		},
		{
			Method:  http.MethodGet,
			URL:     "/api/v1/user/signed",
			Handler: toHandler(uh.handler.GetPageable),
			Open:    false,
		},
		{
			Method:  http.MethodPost,
			URL:     "/api/v1/user/{id}/update/avatar",
			Handler: toHandler(uh.handler.UpdateAvatar),
			Open:    false,
		},
	}
}

func toHandler(f func(rw http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(f)
}
