package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"kredit-plus/helper"
	"kredit-plus/model/web"
	"net/http"
	"strings"
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
	} else if url == "/api/auth/login" || url == "/api/auth/register" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		ctx := TokenValidationMiddleware(writer, request)
		if ctx != nil {
			middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))
		} else {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToReponseBody(writer, webResponse)
		}
	}

}

func TokenValidationMiddleware(w http.ResponseWriter, r *http.Request) context.Context {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println(tokenString, "isi token")

	claims := &helper.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return helper.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	fmt.Println("lolos")

	// Set claims in context
	ctx := context.WithValue(r.Context(), "claims", claims)
	return ctx
}
