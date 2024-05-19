package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/moosch/go-api-example/api"
	"github.com/moosch/go-api-example/internal/cache"
	"github.com/moosch/go-api-example/internal/database"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var username string = r.URL.Query().Get("username")
        var token = r.Header.Get("Authorization")
        var err error

        if username == "" || token == "" {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(w, UnAuthorizedError)
            return
        }

        var db *database.DatabaseInterface
        db, err = database.GetDatabaseConnection()
        if err != nil {
            api.InternalErrorHandler(w)
            return
        }

        // Try the cache first
        cachedToken, found := cache.AuthCache.Get(username)
        if found {
            if token == cachedToken {
                fmt.Println("Cache hit")
                next.ServeHTTP(w, r)
                return
            }
            fmt.Println("Cache hit but invalid")
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(w, UnAuthorizedError)
            return
        }
        fmt.Println("Cache miss")

        var loginDetails *database.LoginDetails
        loginDetails = (*db).GetLoginDetails(username)

        if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(w, UnAuthorizedError)
            return
        }

        next.ServeHTTP(w, r)
    })
}

