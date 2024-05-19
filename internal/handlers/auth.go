package handlers

import (
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	"github.com/moosch/go-api-example/api"
	"github.com/moosch/go-api-example/internal/cache"
	"github.com/moosch/go-api-example/internal/database"
	"github.com/moosch/go-api-example/internal/tokens"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var decoder = schema.NewDecoder()
    var params api.LoginParams
    var err error

    err = decoder.Decode(&params, r.URL.Query())
    if err != nil {
        log.Error(err)
        api.RequestErrorHandler(w, err)
        return
    }

    var db *database.DatabaseInterface
    db, err = database.GetDatabaseConnection()
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var loginDetails *database.LoginDetails
    loginDetails = (*db).GetLoginDetails(params.Username)
    if loginDetails == nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    // Update cache
    token, err := tokens.GenerateToken()
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
    cache.AuthCache.Set((*loginDetails).Username, token)
    go func() {
        (*db).UpdateAuthToken(loginDetails.Username, token)
    }()

    sendResponse(w, api.LoginResponse{
        Username:  (*loginDetails).Username,
        AuthToken: token,
    })
}

func Logout(w http.ResponseWriter, r *http.Request) {
}

