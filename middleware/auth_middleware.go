package middleware

import (
	"database/sql"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
	DB      *sql.DB
}

func NewAuthMiddleware(handler http.Handler, db *sql.DB) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, DB: db}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.String()

	if url == "/api/installment-process" {
		db, err := middleware.DB.Begin()
		helper.PanicIfError(err)
		apiKey := request.Header.Get("X-Api-Key")
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM merchants WHERE api_key=$1)", apiKey).Scan(&exists)
		helper.PanicIfError(err)

		if exists {
			//ok
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			//error
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToReponseBody(writer, webResponse)
		}
	} else {
		middleware.Handler.ServeHTTP(writer, request)
	}

}
