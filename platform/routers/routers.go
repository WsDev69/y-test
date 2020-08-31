package routers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"y-test/internal/glue/middleware"
)

type Routers interface {
	Serve()
}

type Router struct {
	Method  string
	URL     string
	Handler http.Handler
	Open    bool
}

type CheckAuth func(http.Handler) http.Handler

type routing struct {
	host          string
	handlers      []*Router
	checkAuthFunc CheckAuth
}

// Initialize is for initialize the handler
func Initialize(host string, routers []*Router, checkAuth CheckAuth) Routers {
	return &routing{host: host, handlers: routers, checkAuthFunc: checkAuth}
}

// Serve is to start serving the HTTP Listener
func (us *routing) Serve() {
	log := logrus.WithContext(context.Background())
	r := mux.NewRouter()
	for _, router := range us.handlers {
		var f = middleware.ContentTypeJson(router.Handler)
		if !router.Open {
			f = us.checkAuthFunc(f)
		}
		r.Handle(router.URL, f).Methods(router.Method)
		log.WithFields(logrus.Fields{
			"method": router.Method,
			"url":    router.URL,
		}).Trace("routing")
	}

	//to get port from cfg
	log.WithField("port", us.host).Info("ready for ListenAndServe")
	err := http.ListenAndServe(us.host, r)
	if err != nil {
		log.WithField("routing", us.host).Fatal(err)
	}
}
