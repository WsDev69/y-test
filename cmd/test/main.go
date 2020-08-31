package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"y-test/internal/glue/routing"
	"y-test/internal/handler/rest"
	"y-test/internal/module/config/env"
	"y-test/internal/module/user"
	"y-test/internal/security"
	"y-test/internal/storage/persistence"
	"y-test/platform/minio"
	"y-test/platform/routers"
	"y-test/platform/sqlite"
)

func init() {
	//todo move to log cong
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

var config = env.NewEnvConfig()

func main() {
	val := config.GetValues()
	if val == nil {
		panic("Configs don't set")
	}

	//Auth Initialization
	auth := security.NewJwtSecurity(val.JWT.Secret)
	authHandler := rest.NewJwtAuth(auth)

	//Database Initialization
	database := persistence.NewUserPersistence(sqlite.InitConnections(val.Sqlite.Path).Open())

	objStoConf := val.Storage
	us := user.NewUserService(database, minio.InitConnection(
		objStoConf.Endpoint,
		objStoConf.AccessKeyId,
		objStoConf.SecretAccessKey,
		objStoConf.UseSSL).Open())

	//Handler And Router
	handler := rest.NewUserHandler(us, auth)
	router := routing.NewUserHandlers(handler).Routers()

	//Initialized And Start Serve
	servant := routers.Initialize(fmt.Sprintf(":%d", val.Server.Port), router, authHandler.Authentication)
	go servant.Serve()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Info("Application gracefully terminated")
}
