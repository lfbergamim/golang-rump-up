package middleware

import (
	"log"
	"net/http"

	"github.com/lbergamim-daitan/golang-rump-up/internal/auth"
	"github.com/lbergamim-daitan/golang-rump-up/internal/config"
	"github.com/lbergamim-daitan/golang-rump-up/internal/models"
	"github.com/lbergamim-daitan/golang-rump-up/internal/mysql"
	"github.com/lbergamim-daitan/golang-rump-up/internal/orm"
	"github.com/lbergamim-daitan/golang-rump-up/internal/responses"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		nextFunc(w, r)
	}
}

func DatabaseChoose() models.DatabaseInterface {
	if config.DBImplem == "mysql" {
		return &mysql.Mysql{}
	}
	return &orm.ORM{}
}
