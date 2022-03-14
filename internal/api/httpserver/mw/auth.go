package mw

import (
	"net/http"
	"ted/internal/api/auth"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func Auth(auth *auth.AuthService, lg *log.Entry) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			r := render.New()
			userID := request.Header.Get("X-UserId")
			userKey, err := auth.GetUserKey(request.Context(), userID)
			if err != nil {
				lg.Error(err)
				err = r.JSON(writer, http.StatusInternalServerError, map[string]string{"Error": "Internal Server Error"})
				if err != nil {
					lg.Error(err)
					return
				}
				return
			}
			digest := request.Header.Get("X-Digest")

			ok, err := auth.CheckDigest(digest, userKey, request)
			if err != nil {
				lg.Error(err)
				err = r.JSON(writer, http.StatusUnauthorized, map[string]string{"Error": "Unauthorized"})
				if err != nil {
					lg.Error(err)
					return
				}
				return
			}
			if !ok {
				err = r.JSON(writer, http.StatusUnauthorized, map[string]string{"Error": "Unauthorized"})
				if err != nil {
					lg.Error(err)
					return
				}
			}
		})
	}
}
