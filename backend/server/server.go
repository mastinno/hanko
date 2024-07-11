package server

import (
	"github.com/labstack/echo/v4"
	"github.com/teamhanko/hanko/backend/config"
	"github.com/teamhanko/hanko/backend/handler"
	"github.com/teamhanko/hanko/backend/mapper"
	"github.com/teamhanko/hanko/backend/persistence"
	"sync"
)

func StartPublic(cfg *config.Config, wg *sync.WaitGroup, persister persistence.Persister, prometheus echo.MiddlewareFunc, authenticatorMetadata mapper.AuthenticatorMetadata) {
	defer wg.Done()
	router := handler.NewPublicRouter(cfg, persister, prometheus, authenticatorMetadata)
	//router.Logger.Fatal(router.Start(cfg.Server.Public.Address))
	router.Logger.Fatal(router.StartTLS(cfg.Server.Public.Address, "/etc/config/keys/servercrt.pem", "/etc/config/keys/serverkey.pem"))
}

func StartAdmin(cfg *config.Config, wg *sync.WaitGroup, persister persistence.Persister, prometheus echo.MiddlewareFunc) {
	defer wg.Done()
	router := handler.NewAdminRouter(cfg, persister, prometheus)
	//router.Logger.Fatal(router.Start(cfg.Server.Admin.Address))
	router.Logger.Fatal(router.StartTLS(cfg.Server.Admin.Address, "/etc/config/keys/servercrt.pem", "/etc/config/keys/serverkey.pem"))
}
