// DO NOT EDIT, generated by goctl
package handler

import (
	"net/http"

	open "jwttoken/internal/handler/open"
	user "jwttoken/internal/handler/user"
	"jwttoken/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.Use(Auth)
	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/open/authorization",
			Handler: open.AuthorizationHandler(serverCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/user/update",
			Handler: user.EdituserHandler(serverCtx),
		},
	})
}