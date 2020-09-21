package handler

import (
	"net/http"

	"jwttoken/internal/logic/user"
	"jwttoken/internal/svc"
	"jwttoken/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func EdituserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEdituserLogic(r.Context(), ctx)
		resp, err := l.Edituser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
