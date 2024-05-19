package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/moosch/go-api-example/api"
	"github.com/moosch/go-api-example/internal/database"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
    var params api.GetAccountParams
    var err error
    if err = parseRequestUrlQuery(r.URL.Query(), &params); err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var db *database.DatabaseInterface
    db, err = database.GetDatabaseConnection()
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var accountDetails *database.AccountDetails
    accountDetails = (*db).GetAccountDetails(params.Username)
    if accountDetails == nil {
        log.Error(err)
        api.NotFoundErrorHandler(w, "User not found.")
        return
    }

    sendResponse[api.GetAccountResponse](w, api.GetAccountResponse{
        Username: (*accountDetails).Username,
        Age:      (*&accountDetails).Age,
    })
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
    var urlParams api.UpdateAccountParams
    var err error
    if err = parseRequestUrlQuery(r.URL.Query(), &urlParams); err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var bodyParams map[string]interface{}
    if err = json.NewDecoder(r.Body).Decode(&bodyParams); err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
    fmt.Printf("Body: %+v\n", bodyParams)

    var db *database.DatabaseInterface
    db, err = database.GetDatabaseConnection()
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var accountDetails *database.AccountDetails
    accountDetails = (*db).GetAccountDetails(urlParams.Username)
    if accountDetails == nil {
        log.Error(err)
        api.NotFoundErrorHandler(w, "User not found.")
        return
    }

    var newAccountDetails *database.AccountDetails
    newAccountDetails = accountDetails
    if bodyParams["username"] != nil && urlParams.Username != bodyParams["username"] {
        newAccountDetails.Username = bodyParams["username"].(string)
    }
    if bodyParams["age"] != nil {
        newAccountDetails.Age = int(bodyParams["age"].(float64))
    }

    accountDetails, err = (*db).UpdateAccountDetails(urlParams.Username, *newAccountDetails)
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
    sendResponse[api.UpdateAccountResponse](w, api.UpdateAccountResponse{
        Username: (*accountDetails).Username,
        Age:      (*&accountDetails).Age,
    })
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
}

